package repositories

import "portto-assignment/config"

var memeCoinRepository *MemeCoinRepository

func Init(connectionPool *config.DBPool) {
	if memeCoinRepository == nil {
		memeCoinRepository = &MemeCoinRepository{
			pool: *connectionPool, // Store the pointer instead of dereferencing it
		}
	}
}

func GetMemeCoinRepository() *MemeCoinRepository {
	return memeCoinRepository
}
