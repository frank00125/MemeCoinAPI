package services

import (
	"portto-assignment/internal/repositories"
)

func NewMemeCoinService(memeCoinRepository repositories.MemeCoinRepositoryInterface, redisRepository repositories.RedisRepositoryInterface) *MemeCoinService {
	return &MemeCoinService{
		repo:  memeCoinRepository,
		redis: redisRepository,
	}
}

func (service *MemeCoinService) CreateMemeCoin(input CreateMemeCoinInput) (*repositories.MemeCoin, error) {
	return service.repo.CreateOne(input.Name, input.Description)
}

func (service *MemeCoinService) GetMemeCoin(id int) (*repositories.MemeCoin, error) {
	return service.repo.FindOne(id)
}

func (service *MemeCoinService) UpdateMemeCoin(id int, description string) (*repositories.MemeCoin, error) {
	return service.repo.UpdateOne(id, description)
}

func (service *MemeCoinService) DeleteMemeCoin(id int) (*repositories.MemeCoin, error) {
	// Delete popularity_score at redis
	err := service.redis.RemovePopularityScore(id)
	if err != nil {
		return nil, err
	}

	deletedMemeCoin, err := service.repo.DeleteOne(id)
	if err != nil {
		return nil, err
	}

	return deletedMemeCoin, nil
}

func (service *MemeCoinService) PokeMemeCoin(id int) error {
	// Increment popularity_score at redis
	return service.redis.IncrementPopularityScore(id)
}
