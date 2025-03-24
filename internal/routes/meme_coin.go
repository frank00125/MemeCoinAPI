package routes

import (
	"github.com/gin-gonic/gin"

	"portto-assignment/internal/handlers"
)

func SetupMemeCoinRoutes(rg *gin.RouterGroup, handlers handlers.MemeCoinHandlerInterface) {
	memeCoinService := rg.Group("/meme-coin")
	{
		memeCoinService.POST("/create", handlers.CreateMemeCoin)
		memeCoinService.GET("/:id", handlers.GetMemeCoin)
		memeCoinService.PATCH("/:id", handlers.UpdateMemeCoin)
		memeCoinService.DELETE("/:id", handlers.DeleteMemeCoin)
		memeCoinService.POST("/:id/poke", handlers.PokeMemeCoin)
	}
}
