package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/giovanni-liboni/exercise-rest-api-shop/entities"
	"github.com/giovanni-liboni/exercise-rest-api-shop/services"
	log "github.com/sirupsen/logrus"
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
			Message:    "Internal server error",
			Data:       nil,
		}
	}
	return &entities.AppResult{
		Data:       items,
		Message:    "Get all items",
		StatusCode: http.StatusOK,
		Err:        nil,
	}
}

func (i itemHandler) GetItem(ctx *gin.Context) *entities.AppResult {
	itemId := ctx.Param("id")

	// Convert to int64
	itemIdInt, err := strconv.ParseInt(itemId, 10, 64)
	if err != nil {
		log.Errorf("Error converting itemId to int64: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad request",
			Data:       nil,
		}
	}
	item, err := i.itemService.GetItem(ctx, itemIdInt)
	if err != nil {
		log.Errorf("Error getting item: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		}
	}
	return &entities.AppResult{
		Data:       item,
		Message:    "Get item",
		StatusCode: http.StatusOK,
		Err:        nil,
	}
}

func (i itemHandler) CreateItem(ctx *gin.Context) *entities.AppResult {
	item := entities.Item{}
	err := ctx.BindJSON(&item)
	if err != nil {
		log.Errorf("Error binding json: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad request",
			Data:       nil,
		}
	}
	err = i.itemService.CreateItem(ctx, &item)
	if err != nil {
		log.Errorf("Error creating item: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		}
	}
	return &entities.AppResult{
		Data:       nil,
		Message:    "Created item",
		StatusCode: http.StatusCreated,
		Err:        nil,
	}
}

func (i itemHandler) UpdateItem(ctx *gin.Context) *entities.AppResult {
	item := entities.Item{}
	err := ctx.BindJSON(&item)
	if err != nil {
		return &entities.AppResult{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad request",
			Data:       nil,
		}
	}
	err = i.itemService.UpdateItem(ctx, &item)
	if err != nil {
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		}
	}
	return &entities.AppResult{
		Data:       nil,
		Message:    "Updated item",
		StatusCode: http.StatusOK,
		Err:        nil,
	}
}

func (i itemHandler) DeleteItem(ctx *gin.Context) *entities.AppResult {
	itemId := ctx.Param("id")

	// Convert to int64
	itemIdInt, err := strconv.ParseInt(itemId, 10, 64)
	if err != nil {
		return &entities.AppResult{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad request",
			Data:       nil,
		}
	}

	err = i.itemService.DeleteItem(ctx, itemIdInt)
	if err != nil {
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		}
	}
	return &entities.AppResult{
		Data:       nil,
		Message:    "Deleted item",
		StatusCode: http.StatusOK,
		Err:        nil,
	}
}

func (i itemHandler) PurchaseItem(ctx *gin.Context) *entities.AppResult {
	// Get the authenticated user
	user := ctx.MustGet("userID").(*entities.User)
	itemId := ctx.Param("id")

	// Convert to int64
	itemIdInt, err := strconv.ParseInt(itemId, 10, 64)
	if err != nil {
		log.Errorf("Error converting itemId to int64: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad request",
			Data:       nil,
		}
	}

	// Get the user by username
	order, err := i.itemService.PurchaseItem(ctx, itemIdInt, user.ID)
	// If the item is not found, return a 404
	if err == sql.ErrNoRows {
		return &entities.AppResult{
			StatusCode: http.StatusNotFound,
			Message:    "Item not found",
			Data:       nil,
		}
	}

	if err != nil {
		log.Errorf("Error purchasing item: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		}
	}

	return &entities.AppResult{
		Data:       order,
		Message:    "Purchased item: order created",
		StatusCode: http.StatusCreated,
		Err:        nil,
	}
}
