package main

import (
	"portto-assignment/config"
	"portto-assignment/repositories"
	"portto-assignment/routes"
	"portto-assignment/services"

	"github.com/gin-gonic/gin"
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

	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router)

	router.Run(":8080")

}
