package services

import (
	"errors"
	"portto-assignment/internal/repositories"
)

func NewMemeCoinService(memeCoinRepository repositories.MemeCoinRepositoryInterface) *MemeCoinService {
	return &MemeCoinService{
		repo: memeCoinRepository,
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

func (service *MemeCoinService) PokeMemeCoin(id int) (*repositories.MemeCoin, error) {
	// Poke meme coin
	err := service.repo.PokeOne(id)
	if err != nil {
		return nil, err
	}

	// Get updated meme coin after transaction
	updatedMemeCoin, err := service.repo.FindOne(id)
	if err != nil {
		return nil, err
	}

	return updatedMemeCoin, nil
}
