package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
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
		db:        db,
		redis:     redis,
		config:    config,
		dirtyKeys: make(chan string, config.SyncBatchSize*2), // Buffer size based on batch size
	}

	if config.NeedToSync {
		// Sync Redis with the database
		isDone := make(chan bool)
		go func() {
			repo.setDataToRedis()
			isDone <- true
		}()
		<-isDone

		// Start the sync worker to sync Redis with the database
		repo.startSyncWorker()
	}

	return repo
}

func (r *RedisCachedRepository) IncrementPopularityScore(id int) error {
	// Atomic increment in Redis
	_, err := r.redis.IncrBy(context.Background(),
		fmt.Sprintf("meme:popularity_score:%d", id), 1).Result()
	if err != nil {
		return err
	}

	// Signal that this key needs to be synced
	select {
	case r.dirtyKeys <- strconv.Itoa(id):
		// Key queued for sync
	default:
		// Channel full, key will be synced on next cycle
	}

	return nil
}

func (r *RedisCachedRepository) startSyncWorker() {
	ticker := time.NewTicker(r.config.SyncInterval)
	pendingCounts := 0
	pendingSync := make(map[int]bool) // Just tracking which IDs need sync

	go func() {
		for {
			select {
			case id := <-r.dirtyKeys:
				idNum, _ := strconv.Atoi(id)
				if !pendingSync[idNum] {
					pendingSync[idNum] = true
					pendingCounts++
				}

				// If we have enough pending items, trigger a sync
				if pendingCounts >= r.config.SyncBatchSize {
					r.syncBatch(pendingSync)
					pendingSync = make(map[int]bool)
					pendingCounts = 0
				}

			case <-ticker.C:
				// Time-based sync for any remaining items
				if pendingCounts > 0 {
					r.syncBatch(pendingSync)
					pendingSync = make(map[int]bool)
					pendingCounts = 0
				}
			}

		}
	}()
}

func (r *RedisCachedRepository) syncBatch(ids map[int]bool) {
	// Start a transaction
	ctx := context.Background()
	tx, err := r.db.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		return
	}

	for id := range ids {
		// Get current score from Redis - this will include ALL increments that have happened
		scoreKey := fmt.Sprintf("meme:popularity_score:%d", id)
		score, err := r.redis.Get(context.Background(), scoreKey).Int()
		if err != nil {
			continue
		}

		// Update database with the accurate count from Redis
		_, err = tx.ExecContext(ctx, "UPDATE meme_coins SET popularity_score = $2 WHERE id = $1", id, score)
		if err != nil {
			log.Printf("Error updating score for %d: %v", id, err)
		}
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		log.Printf("Error committing transaction: %v", err)
		tx.Rollback()
	}
}

func (r *RedisCachedRepository) setDataToRedis() {
	// Fetch all popularity scores from the database
	var popularityScoreRows []memeCoinPopularityScore
	const limit = 100
	page := 0
	for {
		rows, err := r.db.Query("SELECT id, popularity_score FROM meme_coins LIMIT $1 OFFSET $2,", limit, limit*page)
		if err != nil {
			log.Fatalf("Error fetching popularity scores: %v\b", err)
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
			log.Fatalf("Error setting popularity scores in Redis: %v\n", err)
			return
		}

		if len(popularityScoreRows) < limit {
			break
		}
		page++
	}

}
