package main

import (
	_ "portto-assignment/api"
	"portto-assignment/config"
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

	// Inject database connection pools
	memeCoinRepository := repositories.NewMemeCoinRepository(connectionPool)

	// Inject repositories
	memeCoinService := services.NewMemeCoinService(memeCoinRepository)

	// Inject services
	memeCoinHandler := handlers.NewMemeCoinHandler(memeCoinService)

	// Setup routes
	router := routes.NewRouter(memeCoinHandler)
	router.Run(":8080")
}
