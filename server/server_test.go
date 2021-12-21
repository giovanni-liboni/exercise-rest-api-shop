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

func TestGetItemRoute(t *testing.T) {
	globalConfig := config.LoadConfig("../.test.env")
	// Initialize the router
	router := SetupRouter(globalConfig)
	r := gofight.New()

	var token string

	// First authenticate the user
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

	// Then get the item
	r.GET("/items/1").
		SetHeader(gofight.H{
			"Authorization": "Bearer " + token,
		}).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
			assert.Equal(t, "The Misty Cup", gjson.Get(r.Body.String(), "data.name").String())
		})

}

func TestPurchaseRoute(t *testing.T) {
	globalConfig := config.LoadConfig("../.test.env")
	// Initialize the router
	router := SetupRouter(globalConfig)
	r := gofight.New()

	var token string

	// First authenticate the user
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

	// Then purchase an item
	r.POST("/items/1/purchase").
		SetHeader(gofight.H{
			"Authorization": "Bearer " + token,
		}).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusCreated, r.Code)
		})
}

func TestPurchaseRoute_ItemNotPresent(t *testing.T) {
	globalConfig := config.LoadConfig("../.test.env")
	// Initialize the router
	router := SetupRouter(globalConfig)
	r := gofight.New()

	var token string

	// First authenticate the user
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

	// Then purchase an item
	r.POST("/items/435/purchase").
		SetHeader(gofight.H{
			"Authorization": "Bearer " + token,
		}).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusNotFound, r.Code)
		})
}

func TestPayOrderRoute(t *testing.T) {
	globalConfig := config.LoadConfig("../.test.env")
	// Initialize the router
	router := SetupRouter(globalConfig)
	r := gofight.New()

	var token string

	// First authenticate the user
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

	// Then pay for the order
	r.POST("/orders/6/pay").
		SetHeader(gofight.H{
			"Authorization": "Bearer " + token,
		}).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestGetUserOrders(t *testing.T) {
	globalConfig := config.LoadConfig("../.test.env")
	// Initialize the router
	router := SetupRouter(globalConfig)
	r := gofight.New()

	var token string

	// First authenticate the user
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

	// Then get the user's orders

	r.GET("/users/me/orders").
		SetHeader(gofight.H{
			"Authorization": "Bearer " + token,
		}).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
			// Get the orders from the response
			orders := gjson.Get(r.Body.String(), "data")
			assert.GreaterOrEqual(t, len(orders.Array()), 3) // The user should have at least 3 orders (one is
																 // created by the test)
		})
}

func TestGetUserOrderItems(t *testing.T) {
	globalConfig := config.LoadConfig("../.test.env")
	// Initialize the router
	router := SetupRouter(globalConfig)
	r := gofight.New()

	var token string

	// First authenticate the user
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

	// Then get the user's order items

	r.GET("/users/me/orders/6/items").
		SetHeader(gofight.H{
			"Authorization": "Bearer " + token,
		}).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
			// Get the orders from the response
			items := gjson.Get(r.Body.String(), "data")
			assert.Equal(t, 1, len(items.Array()))
		})
}

func TestGetUserOrderItems_Forbidden(t *testing.T) {
	globalConfig := config.LoadConfig("../.test.env")
	// Initialize the router
	router := SetupRouter(globalConfig)
	r := gofight.New()

	var token string

	// First authenticate the user
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

	// Then get the user's order items

	r.GET("/users/me/orders/1/items").
		SetHeader(gofight.H{
			"Authorization": "Bearer " + token,
		}).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusForbidden, r.Code)
		})
}

func TestCreateUser_UserAlreadyExists(t *testing.T) {
	globalConfig := config.LoadConfig("../.test.env")
	// Initialize the router
	router := SetupRouter(globalConfig)
	r := gofight.New()

	// Create a new user
	r.POST("/auth/register").
		SetJSON(gofight.D{
			"username": "test",
			"password": "test",
		}).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusConflict, r.Code)
		})
}

func TestPublicStatistics(t *testing.T) {
	globalConfig := config.LoadConfig("../.test.env")
	// Initialize the router
	router := SetupRouter(globalConfig)
	r := gofight.New()

	// Get the public statistics
	r.GET("/statistics").
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestAdminStatistics(t *testing.T) {
	globalConfig := config.LoadConfig("../.test.env")
	// Initialize the router
	router := SetupRouter(globalConfig)
	r := gofight.New()

	var token string

	// First authenticate the user
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

	// Then get the user's order items

	r.GET("/admin/statistics").
		SetHeader(gofight.H{
			"Authorization": "Bearer " + token,
		}).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}