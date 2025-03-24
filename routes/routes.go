package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		SetupMemeCoinRoutes(v1)
		SetupDocsRoutes(v1)
	}

	return router
}
