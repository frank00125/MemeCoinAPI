package handlers

import (
	"fmt"
	"net/http"
	"portto-assignment/services"

	"github.com/gin-gonic/gin"
)

func CreateMemeCoinHandler(context *gin.Context) {
	// Get request body
	var reqBody *struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description" binding:"-"`
	}
	err := context.BindJSON(&reqBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	// Call service
	name := reqBody.Name
	description := reqBody.Description
	memeCoinService := services.GetMemeCoinService()
	newMemeCoin, err := memeCoinService.CreateMemeCoin(services.CreateMemeCoinInput{
		Name:        name,
		Description: description,
	})
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database Error",
			"error":   err.Error(),
		})
		return
	}

	if newMemeCoin == nil {
		context.JSON(http.StatusConflict, gin.H{
			"message": "MemeCoin already exists",
		})
		return
	}

	context.JSON(http.StatusOK, newMemeCoin)
}

func GetMemeCoinHandler(context *gin.Context) {
	var reqBody *struct {
		Id int `uri:"id" binding:"required"`
	}
	err := context.BindUri(&reqBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid MemeCoin ID",
			"error":   err.Error(),
		})
		return
	}

	id := reqBody.Id
	memeCoinService := services.GetMemeCoinService()
	memeCoin, err := memeCoinService.GetMemeCoin(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database Error",
			"error":   err.Error(),
		})
		return
	}

	if memeCoin == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "MemeCoin not found",
		})
		return
	}

	context.JSON(http.StatusOK, memeCoin)
}

func UpdateMemeCoinHandler(context *gin.Context) {
	var urlParams *struct {
		Id int `uri:"id" binding:"required"`
	}
	// from URL
	err := context.ShouldBindUri(&urlParams)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid MemeCoin ID",
			"error":   err.Error(),
		})
		return
	}

	// from body
	var reqBody *struct {
		Description string `json:"description" binding:"required"`
	}
	err = context.ShouldBindJSON(&reqBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	memeCoinService := services.GetMemeCoinService()
	updatedMemeCoin, err := memeCoinService.UpdateMemeCoin(urlParams.Id, reqBody.Description)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database Error",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, updatedMemeCoin)
}

func DeleteMemeCoinHandler(context *gin.Context) {
	var reqBody *struct {
		Id int `uri:"id" binding:"required"`
	}
	err := context.BindUri(&reqBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid MemeCoin ID",
			"error":   err.Error(),
		})
		return
	}

	id := reqBody.Id
	memeCoinService := services.GetMemeCoinService()
	deletedMemeCoin, err := memeCoinService.DeleteMemeCoin(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database Error",
			"error":   err.Error(),
		})
		return
	}

	if deletedMemeCoin == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "MemeCoin not found",
		})
		return
	}

	context.JSON(http.StatusOK, deletedMemeCoin)
}

func PokeMemeCoinHandler(context *gin.Context) {
	var reqBody *struct {
		Id int `uri:"id" binding:"required"`
	}
	err := context.BindUri(&reqBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid MemeCoin ID",
			"error":   err.Error(),
		})
		return
	}

	id := reqBody.Id
	memeCoinService := services.GetMemeCoinService()
	pokedMemeCoin, err := memeCoinService.PokeMemeCoin(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database Error",
			"error":   err.Error(),
		})
		return
	}

	if pokedMemeCoin == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "MemeCoin not found",
		})
		return
	}

	context.JSON(http.StatusOK, pokedMemeCoin)
}
