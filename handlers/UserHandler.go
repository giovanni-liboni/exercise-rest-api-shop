package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/giovanni-liboni/exercise-rest-api-shop/entities"
	"github.com/giovanni-liboni/exercise-rest-api-shop/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type userHandler struct {
	userService services.UserService
	orderService services.OrderService
}

type UserHandler interface {
	GetUserOrders(ctx *gin.Context) *entities.AppResult
	GetUserOrderItems(ctx *gin.Context) *entities.AppResult
	CreateUser(ctx *gin.Context) *entities.AppResult
}

func InitUserHandler(userService services.UserService, orderService services.OrderService) UserHandler {
	return &userHandler{
		userService: userService,
		orderService: orderService,
	}
}

func (u userHandler) GetUserOrders(ctx *gin.Context) *entities.AppResult {
	// Get the current authenticated userID
	user := ctx.MustGet("userID").(*entities.User)

	// Get the orders for the current user
	orders, err := u.orderService.GetOrdersByUser(ctx, user.ID)
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

func (u userHandler) GetUserOrderItems(ctx *gin.Context) *entities.AppResult {
	// Get the current authenticated userID
	user := ctx.MustGet("userID").(*entities.User)

	// Get the orderID
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

	// Get the orders for the current user
	order, err := u.orderService.GetOrder(ctx, orderIdInt)
	if err != nil {
		log.Errorf("Error getting orders: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		}
	}

	// Check the current user is the owner of the order
	if order.UserID != user.ID {
		log.Errorf("User %v is not the owner of order %v", user.ID, order.ID)
		return &entities.AppResult{
			StatusCode: http.StatusForbidden,
			Message:    "Forbidden",
			Data:       nil,
		}
	}

	return &entities.AppResult{
		StatusCode: http.StatusOK,
		Message:    "OK",
		Data:       order.Items,
	}
}

func (u userHandler) CreateUser(ctx *gin.Context) *entities.AppResult {
	// Sanitize the request body
	var user entities.User
	err := ctx.BindJSON(&user)
	if err != nil {
		log.Errorf("Error binding JSON: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad request",
			Data:       nil,
		}
	}

	// Create the user
	err = u.userService.CreateUser(ctx, &user)
	// If the user already exists, return a conflict
	if err == entities.ErrUserAlreadyExists {
		log.Errorf("User already exists: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusConflict,
			Message:    "Conflict",
			Data:       nil,
		}
	}

	if err != nil {
		log.Errorf("Error creating user: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		}
	}

	return &entities.AppResult{
		StatusCode: http.StatusCreated,
		Message:    "Created",
		Data:       user,
	}

}