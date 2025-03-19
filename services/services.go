package services

import "portto-assignment/repositories"

var memeCoinService *MemeCoinService

func Init(memeCoinRepository repositories.MemeCoinRepositoryInterface) {
	memeCoinService = &MemeCoinService{
		repo: memeCoinRepository,
	}
}

func GetMemeCoinService() *MemeCoinService {
	return memeCoinService
}
