package main

import (
	"portto-assignment/config"
	"portto-assignment/repositories"
	"portto-assignment/routes"
	"portto-assignment/services"
)

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
