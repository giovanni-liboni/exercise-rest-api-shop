package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/giovanni-liboni/exercise-rest-api-shop/entities"
	"github.com/giovanni-liboni/exercise-rest-api-shop/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type orderHandler struct {
	orderService services.OrderService
	userService  services.UserService
}

type OrderHandler interface {
	GetOrder(ctx *gin.Context) *entities.AppResult
	GetOrders(ctx *gin.Context) *entities.AppResult
	PayOrder(ctx *gin.Context) *entities.AppResult
}

func InitOrderHandler(orderService services.OrderService, userService services.UserService) OrderHandler {
	return &orderHandler{
		orderService: orderService,
		userService:  userService,
	}
}

func (o orderHandler) GetOrder(ctx *gin.Context) *entities.AppResult {
	orderID := ctx.Param("id")

	// Convert to int64
	orderIdInt, err := strconv.ParseInt(orderID, 10, 64)
	if err != nil {
		log.Errorf("Error converting itemId to int64: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad request",
			Data:       nil,
		}
	}

	order, err := o.orderService.GetOrder(ctx, orderIdInt)
	if err != nil {
		log.Errorf("Error getting order: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		}
	}

	return &entities.AppResult{
		StatusCode: http.StatusOK,
		Message:    "OK",
		Data:       order,
	}

}

func (o orderHandler) GetOrders(ctx *gin.Context) *entities.AppResult {
	// Get the current authenticated user
	status := ctx.Query("status")
	userID := ctx.MustGet("userID").(int64)

	var orders []*entities.Order
	var err error

	if status != "" {
		orders, err = o.orderService.GetOrdersByUserAndStatus(ctx, userID, status)
	} else {
		orders, err = o.orderService.GetOrdersByUser(ctx, userID)
	}

	if err != nil {
		log.Errorf("Error getting orders: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		}
	}

	return &entities.AppResult{
		StatusCode: http.StatusOK,
		Message:    "OK",
		Data:       orders,
	}
}

func (o orderHandler) PayOrder(ctx *gin.Context) *entities.AppResult {
	// Get the current authenticated user
	user := ctx.MustGet("userID").(*entities.User)
	orderID := ctx.Param("id")

	// Convert to int64
	orderIdInt, err := strconv.ParseInt(orderID, 10, 64)
	if err != nil {
		log.Errorf("Error converting itemId to int64: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad request",
			Data:       nil,
		}
	}

	user, err = o.userService.GetUserByID(ctx, user.ID)
	if err != nil {
		log.Errorf("Error getting user: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		}
	}

	order, err := o.orderService.PayOrder(ctx, orderIdInt, user)
	if err != nil {
		log.Errorf("Error paying order: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		}
	}

	return &entities.AppResult{
		StatusCode: http.StatusOK,
		Message:    "OK",
		Data:       order,
	}
}
