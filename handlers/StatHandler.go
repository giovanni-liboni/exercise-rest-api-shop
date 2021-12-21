package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/giovanni-liboni/exercise-rest-api-shop/entities"
	"github.com/giovanni-liboni/exercise-rest-api-shop/services"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type statHandler struct {
	statService services.StatService
}

type StatHandler interface {
	GetPublicStatistics(ctx *gin.Context) *entities.AppResult
	GetAdminStatistics(ctx *gin.Context) *entities.AppResult
}

func InitStatHandler(statService services.StatService) StatHandler {
	return &statHandler{
		statService: statService,
	}
}

func (s statHandler) GetPublicStatistics(ctx *gin.Context) *entities.AppResult {
	// Retrieve statistics
	statistics, err := s.statService.GetStats(ctx)
	if err != nil {
		log.Errorf("Error getting statistics: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		}
	}

	// Get the total numbers of items
	totalItems, err := s.statService.GetTotalNumberOfItems(ctx)
	if err != nil {
		log.Errorf("Error getting total items: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		}
	}

	// Get the total numbers of users
	totalUsers, err := s.statService.GetTotalNumberOfUsers(ctx)
	if err != nil {
		log.Errorf("Error getting total users: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		}
	}

	// Get the total numbers of orders
	totalOrders, err := s.statService.GetTotalNumberOfOrders(ctx)
	if err != nil {
		log.Errorf("Error getting total orders: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		}
	}

	// Create a new JSON with the three statistics: last month, last week, last day
	return &entities.AppResult{
		StatusCode: http.StatusOK,
		Message:    "Statistics retrieved",
		Data:       gin.H{
			"last_month": statistics[0],
			"last_week": statistics[1],
			"last_day": statistics[2],
			"total_items": totalItems,
			"total_users": totalUsers,
			"total_orders": totalOrders,
		},
	}
}

func (s statHandler) GetAdminStatistics(ctx *gin.Context) *entities.AppResult {
	// Retrieve statistics
	statistics, err := s.statService.GetStats(ctx)
	if err != nil {
		log.Errorf("Error getting statistics: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		}
	}
	// Get the total numbers of items
	totalItems, err := s.statService.GetTotalNumberOfItems(ctx)
	if err != nil {
		log.Errorf("Error getting total items: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		}
	}

	// Get the total numbers of users
	totalUsers, err := s.statService.GetTotalNumberOfUsers(ctx)
	if err != nil {
		log.Errorf("Error getting total users: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		}
	}

	// Get the total numbers of orders
	totalOrders, err := s.statService.GetTotalNumberOfOrders(ctx)
	if err != nil {
		log.Errorf("Error getting total orders: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		}
	}

	// Get users who spent more
	users, err := s.statService.GetUsersWhoSpendMore(ctx)
	if err != nil {
		log.Errorf("Error getting users who spend more: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		}
	}

	// Get most ordered items
	items, err := s.statService.GetMostOrderedItems(ctx)
	if err != nil {
		log.Errorf("Error getting most ordered items: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		}
	}

	// Get least ordered items
	leastItems, err := s.statService.GetLeastOrderedItems(ctx)
	if err != nil {
		log.Errorf("Error getting least ordered items: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		}
	}

	// Get items not ordered
	itemsNotOrdered, err := s.statService.GetItemsNotOrdered(ctx)
	if err != nil {
		log.Errorf("Error getting items not ordered: %v", err)
		return &entities.AppResult{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		}
	}

	// Create a new JSON with the three statistics: last month, last week, last day
	//
	return &entities.AppResult{
		StatusCode: http.StatusOK,
		Message:    "Statistics retrieved",
		Data:       gin.H{
			"last_month": statistics[0],
			"last_week": statistics[1],
			"last_day": statistics[2],
			"total_items": totalItems,
			"total_users": totalUsers,
			"total_orders": totalOrders,
			"users_spend_more": users,
			"most_ordered_items": items,
			"least_ordered_items": leastItems,
			"items_not_ordered": itemsNotOrdered,
		},
	}
}