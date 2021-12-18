package server

import (
	"github.com/giovanni-liboni/exercise-rest-api-shop/config"
	"github.com/gin-gonic/gin"
)

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


func initRoutes(router *gin.Engine, hds *Handlers, mds *Middlewares) {

}