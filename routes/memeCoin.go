package routes

import (
	"github.com/gin-gonic/gin"

	"portto-assignment/handlers"
)

func SetupMemeCoinRoutes(rg *gin.RouterGroup) {
	memeCoinService := rg.Group("/meme-coin")
	{
		memeCoinService.POST("/create", handlers.CreateMemeCoinHandler)
		memeCoinService.GET("/:id", handlers.GetMemeCoinHandler)
		memeCoinService.PATCH("/:id", handlers.UpdateMemeCoinHandler)
		memeCoinService.DELETE("/:id", handlers.DeleteMemeCoinHandler)
		memeCoinService.POST("/:id/poke", handlers.PokeMemeCoinHandler)
	}
}
