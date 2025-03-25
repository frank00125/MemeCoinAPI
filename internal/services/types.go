package services

import "portto-assignment/internal/repositories"

type MemeCoinService struct {
	repo  repositories.MemeCoinRepositoryInterface
	redis repositories.RedisRepositoryInterface
}

type CreateMemeCoinInput struct {
	Name        string
	Description string
}

type MemeCoinServiceInterface interface {
	CreateMemeCoin(input CreateMemeCoinInput) (*repositories.MemeCoin, error)
	GetMemeCoin(id int) (*repositories.MemeCoin, error)
	UpdateMemeCoin(id int, description string) (*repositories.MemeCoin, error)
	DeleteMemeCoin(id int) (*repositories.MemeCoin, error)
	PokeMemeCoin(id int) error
}
