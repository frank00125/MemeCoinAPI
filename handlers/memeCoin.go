package handlers

import (
	"fmt"
	"net/http"
	"portto-assignment/services"

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

// CreateMemeCoin godoc
//
//	@Summary	Create a MemeCoin
//	@Tags		MemeCoin
//	@Accept		json
//	@Produce	json
//	@Param		body   body handlers.CreateMemeCoinRequestBody true "Request body"
//	@Success	200			{object}	repositories.MemeCoin
//	@Failure	400			{object}	handlers.HttpError
//	@Failure	404			{object}	handlers.HttpError
//	@Failure	409			{object}	handlers.HttpError
//	@Failure	500			{object}	handlers.HttpError
//	@Router		/create [post]
func CreateMemeCoinHandler(context *gin.Context) {
	// Get request body
	var reqBody *CreateMemeCoinRequestBody
	err := context.BindJSON(&reqBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, HttpError{
			Message: "Invalid request body",
			Error:   err.Error(),
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
		context.JSON(http.StatusInternalServerError, HttpError{
			Message: "Database Error",
			Error:   err.Error(),
		})
		return
	}

	if newMemeCoin == nil {
		context.JSON(http.StatusConflict, HttpError{
			Message: "MemeCoin already exists",
		})
		return
	}

	context.JSON(http.StatusOK, newMemeCoin)
}

// GetMemeCoin   godoc
//
//	@Summary	Get a MemeCoin
//	@Tags		MemeCoin
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"MemeCoin ID"
//	@Success	200	{object}	repositories.MemeCoin
//	@Failure	400	{object}	handlers.HttpError
//	@Failure	404	{object}	handlers.HttpError
//	@Failure	500	{object}	handlers.HttpError
//	@Router		/{id} [get]
func GetMemeCoinHandler(context *gin.Context) {
	var reqBody *struct {
		Id int `uri:"id" binding:"required"`
	}
	err := context.BindUri(&reqBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, HttpError{
			Message: "Invalid MemeCoin ID",
			Error:   "Wrong ID format",
		})
		return
	}

	id := reqBody.Id
	memeCoinService := services.GetMemeCoinService()
	memeCoin, err := memeCoinService.GetMemeCoin(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, HttpError{
			Message: "Database Error",
			Error:   err.Error(),
		})
		return
	}

	if memeCoin == nil {
		context.JSON(http.StatusNotFound, HttpError{
			Message: "MemeCoin not found",
		})
		return
	}

	context.JSON(http.StatusOK, memeCoin)
}

type UpdateMemeCoinRequestBody struct {
	Description string `json:"description" binding:"required"`
}

// UpdateMemeCoin  godoc
//
//	@Summary	Update a MemeCoin
//	@Tags		MemeCoin
//	@Accept	json
//	@Produce json
//	@Param id path int true	"MemeCoin ID"
//	@Param body body handlers.UpdateMemeCoinRequestBody true "Request body"
//	@Success	200			{object}	repositories.MemeCoin
//	@Failure	400			{object}	handlers.HttpError
//	@Failure	404			{object}	handlers.HttpError
//	@Failure	500			{object}	handlers.HttpError
//	@Router		/{id} [patch]
func UpdateMemeCoinHandler(context *gin.Context) {
	var urlParams *struct {
		Id int `uri:"id" binding:"required"`
	}
	// from URL
	err := context.ShouldBindUri(&urlParams)
	if err != nil {
		context.JSON(http.StatusBadRequest, HttpError{
			Message: "Invalid MemeCoin ID",
			Error:   "Wrong ID format",
		})
		return
	}

	// from body
	var reqBody *UpdateMemeCoinRequestBody
	err = context.ShouldBindJSON(&reqBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, HttpError{
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	memeCoinService := services.GetMemeCoinService()
	updatedMemeCoin, err := memeCoinService.UpdateMemeCoin(urlParams.Id, reqBody.Description)
	if err != nil {
		context.JSON(http.StatusInternalServerError, HttpError{
			Message: "Database Error",
			Error:   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, updatedMemeCoin)
}

// DeleteMemeCoin   godoc
//
//	@Summary	Delete a MemeCoin
//	@Tags		MemeCoin
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"MemeCoin ID"
//	@Success	200	{object}	repositories.MemeCoin
//	@Failure	400	{object}	handlers.HttpError
//	@Failure	404	{object}	handlers.HttpError
//	@Failure	500	{object}	handlers.HttpError
//	@Router		/{id} [delete]
func DeleteMemeCoinHandler(context *gin.Context) {
	var reqBody *struct {
		Id int `uri:"id" binding:"required"`
	}
	err := context.BindUri(&reqBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, HttpError{
			Message: "Invalid MemeCoin ID",
			Error:   "Wrong ID format",
		})
		return
	}

	id := reqBody.Id
	memeCoinService := services.GetMemeCoinService()
	deletedMemeCoin, err := memeCoinService.DeleteMemeCoin(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, HttpError{
			Message: "Database Error",
			Error:   err.Error(),
		})
		return
	}

	if deletedMemeCoin == nil {
		context.JSON(http.StatusNotFound, HttpError{
			Message: "MemeCoin not found",
		})
		return
	}

	context.JSON(http.StatusOK, deletedMemeCoin)
}

// PokeMemeCoin  godoc
//
//	@Summary	Poke a MemeCoin
//	@Tags		MemeCoin
//	@Accept		json
//	@Produce	json
//	@Param		id			path		int		true	"MemeCoin ID"
//	@Success	200			{object}	repositories.MemeCoin
//	@Failure	400			{object}	handlers.HttpError
//	@Failure	404			{object}	handlers.HttpError
//	@Failure	500			{object}	handlers.HttpError
//	@Router		/{id}/poke [post]
func PokeMemeCoinHandler(context *gin.Context) {
	var reqBody *struct {
		Id int `uri:"id" binding:"required"`
	}
	err := context.BindUri(&reqBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, HttpError{
			Message: "Invalid MemeCoin ID",
			Error:   "Wrong ID format",
		})
		return
	}

	id := reqBody.Id
	memeCoinService := services.GetMemeCoinService()
	pokedMemeCoin, err := memeCoinService.PokeMemeCoin(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, HttpError{
			Message: "Database Error",
			Error:   err.Error(),
		})
		return
	}

	if pokedMemeCoin == nil {
		context.JSON(http.StatusNotFound, HttpError{
			Message: "MemeCoin not found",
		})
		return
	}

	context.JSON(http.StatusOK, pokedMemeCoin)
}
