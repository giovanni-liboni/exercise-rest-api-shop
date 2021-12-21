package main

import (
	"github.com/giovanni-liboni/exercise-rest-api-shop/config"
	"github.com/giovanni-liboni/exercise-rest-api-shop/server"
	"strconv"
)

func main() {
	globalConfig := config.LoadConfig()

	s := server.SetupRouter(globalConfig)
	s.Run(":" + strconv.Itoa(globalConfig.Server.Port))
}
