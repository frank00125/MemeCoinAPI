package services

import (
	"errors"
	"fmt"
	"portto-assignment/internal/repositories"
)

func NewMemeCoinService(memeCoinRepository repositories.MemeCoinRepositoryInterface, redisRepository repositories.RedisRepositoryInterface) *MemeCoinService {
	return &MemeCoinService{
		repo:  memeCoinRepository,
		redis: redisRepository,
	}
}

func (service *MemeCoinService) CreateMemeCoin(input CreateMemeCoinInput) (*repositories.MemeCoin, error) {
	memeCoin, err := service.repo.CreateOne(input.Name, input.Description)
	if err != nil {
		return nil, err
	}
	err = service.redis.Set(service.getMemeCoinPopularityScoreKey(memeCoin.Id), memeCoin.PopularityScore)
	if err != nil {
		return nil, err
	}

	return memeCoin, nil
}

func (service *MemeCoinService) GetMemeCoin(id int) (*repositories.MemeCoin, error) {
	return service.repo.FindOne(id)
}

func (service *MemeCoinService) UpdateMemeCoin(id int, description string) (*repositories.MemeCoin, error) {
	return service.repo.UpdateOne(id, description)
}

func (service *MemeCoinService) DeleteMemeCoin(id int) (*repositories.MemeCoin, error) {
	// Delete popularity_score at redis
	err := service.redis.Delete(service.getMemeCoinPopularityScoreKey(id))
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
	// Check if meme coin exists in Redis
	exist, err := service.redis.Exists(service.getMemeCoinPopularityScoreKey(id))
	if err != nil {
		return err
	}

	if !exist {
		return errors.New("no such meme coin")
	}

	// Increment popularity_score at redis
	return service.redis.IncrBy(service.getMemeCoinPopularityScoreKey(id), 1)
}

func (service *MemeCoinService) getMemeCoinPopularityScoreKey(id int) string {
	return fmt.Sprintf("meme:popularity_score:%d", id)
}
