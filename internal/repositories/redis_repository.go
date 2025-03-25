package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedisCachedRepository(db *sql.DB, redis *redis.Client, config RepositoryConfig) *RedisCachedRepository {
	// Apply defaults if values aren't specified
	if config.SyncBatchSize <= 0 {
		config.SyncBatchSize = DefaultSyncBatchSize // Default value
	}
	if config.SyncInterval <= 0 {
		config.SyncInterval = DefaultSyncInterval // Default value
	}

	repo := &RedisCachedRepository{
		db:       db,
		redis:    redis,
		config:   config,
		syncKeys: make(chan string, config.SyncBatchSize*2), // Buffer size based on batch size
	}

	if config.NeedToSync {
		// Sync Redis with the database
		isDone := make(chan bool)
		go func() {
			repo.setPopularityScoreToRedis()
			isDone <- true
		}()
		<-isDone

		// Start the sync worker to sync Redis with the database
		repo.startPopularityScoreSyncWorker()
	}

	return repo
}

func (r *RedisCachedRepository) IncrBy(key string, increment int) error {
	_, err := r.redis.IncrBy(context.Background(), key, int64(increment)).Result()
	if err != nil {
		return err
	}

	// Add the key to the dirty keys channel
	r.syncKeys <- key

	return nil
}

func (r *RedisCachedRepository) Set(key string, value int) error {
	_, err := r.redis.Set(context.Background(), key, value, 0).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisCachedRepository) Delete(key string) error {
	_, err := r.redis.Del(context.Background(), key).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisCachedRepository) Exists(key string) (bool, error) {
	log.Printf("Checking key: %s", key)
	count, err := r.redis.Exists(context.Background(), key).Result()
	if err != nil {
		return false, err
	}
	log.Printf("Count: %d", count)

	return count > 0, nil
}

func (r *RedisCachedRepository) startPopularityScoreSyncWorker() {
	ticker := time.NewTicker(r.config.SyncInterval)
	pendingCounts := 0
	pendingSync := make(map[string]bool) // Just tracking which IDs need sync

	go func() {
		for {
			select {
			case key := <-r.syncKeys:
				if !pendingSync[key] {
					pendingSync[key] = true
					pendingCounts++
				}

				// If we have enough pending items, trigger a sync
				if pendingCounts >= r.config.SyncBatchSize {
					r.syncPopularityScoreBatch(pendingSync)
					pendingSync = make(map[string]bool)
					pendingCounts = 0
				}

			case <-ticker.C:
				// Time-based sync for any remaining items
				if pendingCounts > 0 {
					r.syncPopularityScoreBatch(pendingSync)
					pendingSync = make(map[string]bool)
					pendingCounts = 0
				}
			}

		}
	}()
}

func (r *RedisCachedRepository) syncPopularityScoreBatch(keysExistMap map[string]bool) {
	// Start a transaction
	ctx := context.Background()
	tx, err := r.db.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		return
	}

	keys := []string{}
	for key := range keysExistMap {
		// Get current score from Redis - this will include ALL increments that have happened
		score, err := r.redis.Get(context.Background(), key).Int()
		if err != nil {
			continue
		}

		// Update database with the accurate count from Redis
		tokens := strings.Split(key, ":")
		id, _ := strconv.Atoi(tokens[len(tokens)-1])
		_, err = tx.ExecContext(ctx, "UPDATE meme_coins SET popularity_score = $2 WHERE id = $1", id, score)
		if err != nil {
			log.Printf("Error updating score for %s: %v", key, err)
		}

		keys = append(keys, key)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		log.Printf("Error committing transaction: %v", err)
		tx.Rollback()
	}

	// Log the sync
	log.Printf("Synced keys: %v", keys)
}

func (r *RedisCachedRepository) setPopularityScoreToRedis() {
	// Fetch all popularity scores from the database
	var popularityScoreRows []memeCoinPopularityScore
	const limit = 100
	page := 0
	for {
		rows, err := r.db.Query("SELECT id, popularity_score FROM meme_coins LIMIT $1 OFFSET $2", limit, limit*page)
		if err != nil {
			log.Printf("Error fetching popularity scores: %v\b", err)
			return
		}

		for rows.Next() {
			var popularityScoreRow memeCoinPopularityScore
			rows.Scan(&popularityScoreRow.Id, &popularityScoreRow.PopularityScore)
			popularityScoreRows = append(popularityScoreRows, popularityScoreRow)
		}

		ctx := context.Background()
		pipe := r.redis.Pipeline()
		for _, row := range popularityScoreRows {
			pipe.Set(ctx, fmt.Sprintf("meme:popularity_score:%d", row.Id), row.PopularityScore, 0)
		}
		_, err = pipe.Exec(ctx)
		if err != nil {
			log.Printf("Error setting popularity scores in Redis: %v\n", err)
			return
		}

		if len(popularityScoreRows) < limit {
			break
		}
		page++
	}

}
