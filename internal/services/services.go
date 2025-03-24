package services

import (
	"errors"
	"portto-assignment/internal/repositories"
)

func NewMemeCoinService(memeCoinRepository repositories.MemeCoinRepositoryInterface, redisRepository repositories.RedisRepositoryInterface) *MemeCoinService {
	return &MemeCoinService{
		repo:  memeCoinRepository,
		redis: redisRepository,
	}
}

func (service *MemeCoinService) CreateMemeCoin(input CreateMemeCoinInput) (*repositories.MemeCoin, error) {
	if input.Name == "" {
		return nil, errors.New("name is required")
	}

	return service.repo.CreateOne(input.Name, input.Description)
}

func (service *MemeCoinService) GetMemeCoin(id int) (*repositories.MemeCoin, error) {

	return service.repo.FindOne(id)
}

func (service *MemeCoinService) UpdateMemeCoin(id int, description string) (*repositories.MemeCoin, error) {
	return service.repo.UpdateOne(id, description)
}

func (service *MemeCoinService) DeleteMemeCoin(id int) (*repositories.MemeCoin, error) {

	return service.repo.DeleteOne(id)
}

func (service *MemeCoinService) PokeMemeCoin(id int) error {
	// Increment popularity_score at redis
	err := service.redis.IncrementPopularity(id)
	if err != nil {
		return err
	}

	return nil
}
