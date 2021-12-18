package server

import (
	"github.com/gin-gonic/gin"
	"github.com/giovanni-liboni/exercise-rest-api-shop/config"
	"github.com/giovanni-liboni/exercise-rest-api-shop/entities"
	"net/http"
)

type appHandler func(ctx *gin.Context) *entities.AppResult

// SetupRouter creates the router and returns it
func SetupRouter(globalConfig *config.Config) *gin.Engine {
	// 1. Connect to the database
	db := config.ConnectDB(globalConfig)

	// 2. Initialize the repositories
	repositories := InitRepositories(db)

	// 3. Initialize the services
	services := InitServices(repositories)

	// 4. Initialize the handlers
	handlers := InitHandlers(services)

	// 5. Initialize the router with the logger and recovery middleware already attached
	router := gin.Default()

	// 6. Initialize the middlewares
	middlewares := InitMiddlewares(globalConfig, repositories)

	// 7. Initialize the routes
	initRoutes(router, handlers, middlewares)

	return router
}

// ServeHTTP wraps the results from the handler into a JSON response
func ServeHTTP(handle appHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result := handle(ctx)
		if result == nil {
			ctx.JSON(http.StatusInternalServerError, entities.Response{
				Success: false,
				Message: "INTERNAL SERVER ERROR",
				Data:    nil,
			})
		}
		if result.Err == nil {
			ctx.JSON(result.StatusCode, entities.Response{
				Success: true,
				Message: result.Message,
				Data:    result.Data,
			})
		} else {
			ctx.JSON(result.StatusCode, entities.Response{
				Success: false,
				Message: result.Err.Error(),
				Data:    result.Data,
			})
		}
	}
}

// initRoutes initializes the routes for the API endpoints
func initRoutes(router *gin.Engine, hds *Handlers, mds *Middlewares) {

	// Public routes
	router.GET("/dashboard") //TODO
	router.GET("/items")     //TODO
	router.GET("/items/:id") //TODO

	// Authenticated routes (user must be logged in)
	router.POST("/items/:id/purchase")       //TODO
	router.GET("/users/me/orders")           //TODO
	router.GET("/users/me/orders/:id/items") //TODO

	// Authentication routes
	router.POST("/auth/login")    //TODO
	router.POST("/auth/logout")   //TODO
	router.POST("/auth/register") //TODO
	router.POST("/auth/refresh")  //TODO

	// Admin routes
	router.GET("/orders/statistics") //TODO
	router.POST("/items")            //TODO
	router.PUT("/items/:id")         //TODO
	router.DELETE("/items/:id")      //TODO

}
