package main

import (
	_ "portto-assignment/api"
	"portto-assignment/config"
	"portto-assignment/database/seeds"
	"portto-assignment/internal/handlers"
	"portto-assignment/internal/repositories"
	"portto-assignment/internal/routes"
	"portto-assignment/internal/services"
)

// @title			MemeCoin API
// @version		1.0
// @description	This is a simple API for MemeCoin

// @host		localhost:8080
// @BasePath	/v1/meme-coin/
func main() {
	// Get database connection pool
	connectionPool, err := config.NewDatabaseConnectionPool()
	if err != nil {
		panic(err)
	}
	defer connectionPool.Close()

	// Get redis connection
	redisClient, err := config.NewRedisClient()
	if err != nil {
		panic(err)
	}
	defer redisClient.Close()

	// Seed database
	seeds.Seeds(connectionPool)

	// Inject database connection pools
	memeCoinRepository := repositories.NewMemeCoinRepository(connectionPool)
	redisRepository := repositories.NewRedisCachedRepository(connectionPool, redisClient, repositories.RepositoryConfig{
		SyncBatchSize: repositories.DefaultSyncBatchSize,
		SyncInterval:  repositories.DefaultSyncInterval,
		NeedToSync:    true,
	})

	// Inject repositories
	memeCoinService := services.NewMemeCoinService(memeCoinRepository, redisRepository)

	// Inject services
	memeCoinHandler := handlers.NewMemeCoinHandler(memeCoinService)

	// Setup routes
	router := routes.NewRouter(memeCoinHandler)
	router.Run(":8080")
}
