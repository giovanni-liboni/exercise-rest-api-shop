package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/giovanni-liboni/exercise-rest-api-shop/entities"
	"github.com/giovanni-liboni/exercise-rest-api-shop/services"
	"net/http"
	"strconv"
)

type itemHandler struct {
	itemService services.ItemService
}

type ItemHandler interface {
	GetAllItems(ctx *gin.Context) *entities.AppResult
	GetItem(ctx *gin.Context) *entities.AppResult
	CreateItem(ctx *gin.Context) *entities.AppResult
	UpdateItem(ctx *gin.Context) *entities.AppResult
	DeleteItem(ctx *gin.Context) *entities.AppResult
	PurchaseItem(ctx *gin.Context) *entities.AppResult
}

func InitItemHandler(itemService services.ItemService) ItemHandler {
	return &itemHandler{
		itemService: itemService,
	}
}

func (i itemHandler) GetAllItems(ctx *gin.Context) *entities.AppResult {
	items, err := i.itemService.GetAllItems(ctx)
	if err != nil {
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:  "Internal server error",
			Data: nil,
		}
	}
	return &entities.AppResult{
		Data: items,
		Message: "Get all items",
		StatusCode: http.StatusOK,
		Err: nil,
	}
}

func (i itemHandler) GetItem(ctx *gin.Context) *entities.AppResult {
	itemId := ctx.Param("itemId")

	itemIdInt, err := strconv.Atoi(itemId)
	if err != nil {
		return &entities.AppResult{
			StatusCode: http.StatusBadRequest,
			Message: "Bad request",
			Data: nil,
		}
	}

	item, err := i.itemService.GetItem(ctx, itemIdInt)
	if err != nil {
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:  "Internal server error",
			Data: nil,
		}
	}
	return &entities.AppResult{
		Data: item,
		Message: "Get item",
		StatusCode: http.StatusOK,
		Err: nil,
	}
}

func (i itemHandler) CreateItem(ctx *gin.Context) *entities.AppResult {
	item := entities.Item{}
	err := ctx.BindJSON(&item)
	if err != nil {
		return &entities.AppResult{
			StatusCode: http.StatusBadRequest,
			Message: "Bad request",
			Data: nil,
		}
	}
	err = i.itemService.CreateItem(ctx, &item)
	if err != nil {
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:  "Internal server error",
			Data: nil,
		}
	}
	return &entities.AppResult{
		Data: nil,
		Message: "Created item",
		StatusCode: http.StatusOK,
		Err: nil,
	}
}

func (i itemHandler) UpdateItem(ctx *gin.Context) *entities.AppResult {
	item := entities.Item{}
	err := ctx.BindJSON(&item)
	if err != nil {
		return &entities.AppResult{
			StatusCode: http.StatusBadRequest,
			Message: "Bad request",
			Data: nil,
		}
	}
	err = i.itemService.UpdateItem(ctx, &item)
	if err != nil {
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:  "Internal server error",
			Data: nil,
		}
	}
	return &entities.AppResult{
		Data: nil,
		Message: "Updated item",
		StatusCode: http.StatusOK,
		Err: nil,
	}
}

func (i itemHandler) DeleteItem(ctx *gin.Context) *entities.AppResult {
	itemId := ctx.Param("itemId")

	itemIdInt, err := strconv.Atoi(itemId)
	if err != nil {
		return &entities.AppResult{
			StatusCode: http.StatusBadRequest,
			Message: "Bad request",
			Data: nil,
		}
	}
	err = i.itemService.DeleteItem(ctx, itemIdInt)
	if err != nil {
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:  "Internal server error",
			Data: nil,
		}
	}
	return &entities.AppResult{
		Data: nil,
		Message: "Deleted item",
		StatusCode: http.StatusOK,
		Err: nil,
	}
}

func (i itemHandler) PurchaseItem(ctx *gin.Context) *entities.AppResult {
	panic("implement me")
}