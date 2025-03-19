package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		MemeCoinRouters(v1)
	}
}
