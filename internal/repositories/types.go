package repositories

import (
	"portto-assignment/config"
	"time"
)

type MemeCoin struct {
	Id              int       `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	PopularityScore int       `json:"popularity_score"`
}

type MemeCoinRepositoryInterface interface {
	FindOne(id int) (*MemeCoin, error)
	CreateOne(name string, description string) (*MemeCoin, error)
	UpdateOne(id int, description string) (*MemeCoin, error)
	DeleteOne(id int) (*MemeCoin, error)
	PokeOne(id int) error
}

type MemeCoinRepository struct {
	pool config.DatabaseConnectionPoolInterface
}
