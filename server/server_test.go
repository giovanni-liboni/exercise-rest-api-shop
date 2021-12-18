package server

import (
	"github.com/appleboy/gofight/v2"
	"github.com/giovanni-liboni/exercise-rest-api-shop/config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestPublicDashboardRoute(t *testing.T) {
	globalConfig := config.LoadConfig("../.test.env")
	// Initialize the router
	router := SetupRouter(globalConfig)

	r := gofight.New()

	r.GET("/dashboard").
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
