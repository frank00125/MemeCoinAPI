package routes

import (
	"portto-assignment/internal/handlers"

	"github.com/gin-gonic/gin"
)

func NewRouter(handlers handlers.MemeCoinHandlerInterface) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		SetupMemeCoinRoutes(v1, handlers)
		SetupDocsRoutes(v1)
	}

	return router
}
