package server

import (
	"github.com/appleboy/gofight/v2"
	"github.com/giovanni-liboni/exercise-rest-api-shop/config"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
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

func TestGetAllItemsRoute(t *testing.T) {
	globalConfig := config.LoadConfig("../.test.env")
	// Initialize the router
	router := SetupRouter(globalConfig)

	r := gofight.New()

	r.GET("/items").
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
			assert.Equal(t, "The Misty Cup", gjson.Get(r.Body.String(), "data.0.name").String())
		})
}

func TestLoginRoute_SimpleAuth(t *testing.T) {
	globalConfig := config.LoadConfig("../.test.env")
	// Initialize the router
	router := SetupRouter(globalConfig)

	r := gofight.New()

	r.POST("/auth/login").
		SetJSON(gofight.D{
			"username": "test",
			"password": "test",
		}).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
			assert.NotEmpty(t, gjson.Get(r.Body.String(), "token").String())
		})
}

func TestLogoutRoute(t *testing.T) {
	globalConfig := config.LoadConfig("../.test.env")
	// Initialize the router
	router := SetupRouter(globalConfig)

	r := gofight.New()

	r.POST("/auth/logout").
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestOrderStatisticsRoute_UserGroup(t *testing.T) {
	globalConfig := config.LoadConfig("../.test.env")
	// Initialize the router
	router := SetupRouter(globalConfig)
	r := gofight.New()

	var token string

	r.POST("/auth/login").
		SetJSON(gofight.D{
			"username": "test",
			"password": "test",
		}).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			// Retrieve the bearer token from the body
			tokenRes := gjson.Get(r.Body.String(), "token")
			token = tokenRes.String()
		})
	r.GET("/orders/statistics").
		SetHeader(gofight.H{
			"Authorization": "Bearer " + token,
		}).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusUnauthorized, r.Code)
		})
}

func TestOrderStatisticsRoute_AdminGroup(t *testing.T) {
	globalConfig := config.LoadConfig("../.test.env")
	// Initialize the router
	router := SetupRouter(globalConfig)
	r := gofight.New()

	var token string

	r.POST("/auth/login").
		SetJSON(gofight.D{
			"username": "admin",
			"password": "admin",
		}).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			// Retrieve the bearer token from the body
			tokenRes := gjson.Get(r.Body.String(), "token")
			token = tokenRes.String()
		})
	r.GET("/orders/statistics").
		SetHeader(gofight.H{
			"Authorization": "Bearer " + token,
		}).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
