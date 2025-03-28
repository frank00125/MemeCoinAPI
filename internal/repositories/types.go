package repositories

import (
	"database/sql"
	"time"

	"github.com/redis/go-redis/v9"
)

type MemeCoin struct {
	Id              int       `db:"id" json:"id"`
	Name            string    `db:"name" json:"name"`
	Description     string    `db:"description" json:"description"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	PopularityScore int       `db:"popularity_score" json:"popularity_score"`
}

type memeCoinPopularityScore struct {
	Id              int `db:"id" json:"id"`
	PopularityScore int `db:"popularity_score" json:"popularity_score"`
}

type MemeCoinRepositoryInterface interface {
	FindOne(id int) (*MemeCoin, error)
	CreateOne(name string, description string) (*MemeCoin, error)
	UpdateOne(id int, description string) (*MemeCoin, error)
	DeleteOne(id int) (*MemeCoin, error)
}

type MemeCoinRepository struct {
	db *sql.DB
}

type RedisRepositoryInterface interface {
	IncrBy(key string, increment int) error
	Set(key string, value int) error
	Delete(key string) error
	Exists(key string) (bool, error)
}

type RedisCachedRepository struct {
	db     *sql.DB
	redis  *redis.Client
	config RepositoryConfig
	// Channel for tracking coins that need syncing
	syncKeys chan string
}

type RepositoryConfig struct {
	SyncBatchSize int
	SyncInterval  time.Duration
	NeedToSync    bool
}

const (
	// DefaultSyncBatchSize is the number of records to sync in one batch
	DefaultSyncBatchSize = 100

	// DefaultSyncInterval is how often to sync cache to database
	DefaultSyncInterval = 5 * time.Second
)
