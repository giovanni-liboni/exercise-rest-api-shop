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
	serveHTTP := ServeHTTP

	// Public routes
	router.GET("/dashboard") //TODO
	router.GET("/items", serveHTTP(hds.ItemHandler.GetAllItems))
	router.GET("/items/:id", serveHTTP(hds.ItemHandler.GetItem)) //TODO

	// Authenticated routes (user must be logged in)
	router.POST("/items/:id/purchase", mds.AuthMiddleware.Middleware.MiddlewareFunc())       //TODO
	router.GET("/users/me/orders", mds.AuthMiddleware.Middleware.MiddlewareFunc())           //TODO
	router.GET("/users/me/orders/:id/items", mds.AuthMiddleware.Middleware.MiddlewareFunc()) //TODO

	// Authentication routes
	router.POST("/auth/login", mds.AuthMiddleware.Middleware.LoginHandler)
	router.POST("/auth/logout", mds.AuthMiddleware.Middleware.LogoutHandler)
	router.POST("/auth/register") //TODO
	router.POST("/auth/refresh", mds.AuthMiddleware.Middleware.RefreshHandler)

	// Admin routes
	router.GET("/orders/statistics", mds.AuthMiddleware.Middleware.MiddlewareFunc(), mds.GroupAuthMiddleware.MiddlewareFunc("admin")) //TODO
	router.POST("/items", serveHTTP(hds.ItemHandler.CreateItem), mds.AuthMiddleware.Middleware.MiddlewareFunc(), mds.GroupAuthMiddleware.MiddlewareFunc("admin"))
	router.PUT("/items/:id", serveHTTP(hds.ItemHandler.UpdateItem), mds.AuthMiddleware.Middleware.MiddlewareFunc(), mds.GroupAuthMiddleware.MiddlewareFunc("admin"))
	router.DELETE("/items/:id", serveHTTP(hds.ItemHandler.DeleteItem), mds.AuthMiddleware.Middleware.MiddlewareFunc(), mds.GroupAuthMiddleware.MiddlewareFunc("admin"))

}
