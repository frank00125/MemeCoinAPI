package main

import (
	"portto-assignment/config"
	_ "portto-assignment/docs"
	"portto-assignment/repositories"
	"portto-assignment/routes"
	"portto-assignment/services"
)

// @title			MemeCoin API
// @version		1.0
// @description	This is a simple API for MemeCoin

// @host		localhost:8080
// @BasePath	/v1/meme-coin/
func main() {
	// Load environment variables
	config.LoadEnvVars()

	// Inject database connection pool
	config.InitDatabase()
	connectionPool := config.GetConnection()

	// Inject repositories
	repositories.Init(connectionPool)
	memeCoinRepository := repositories.GetMemeCoinRepository()

	// Inject services
	services.Init(memeCoinRepository)

	// Setup routes
	router := routes.SetupRouter()
	router.Run(":8080")
}
