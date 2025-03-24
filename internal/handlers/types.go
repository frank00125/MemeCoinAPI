package handlers

import (
	"portto-assignment/internal/services"

	"github.com/gin-gonic/gin"
)

type HttpError struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type CreateMemeCoinRequestBody struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"-"`
}

type UpdateMemeCoinRequestBody struct {
	Description string `json:"description" binding:"required"`
}

type MemeCoinHandlerInterface interface {
	CreateMemeCoin(context *gin.Context)
	GetMemeCoin(context *gin.Context)
	UpdateMemeCoin(context *gin.Context)
	DeleteMemeCoin(context *gin.Context)
	PokeMemeCoin(context *gin.Context)
}

type MemeCoinHandler struct {
	service services.MemeCoinServiceInterface
}
