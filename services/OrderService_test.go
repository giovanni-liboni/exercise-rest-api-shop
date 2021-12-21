package services

import (
	"context"
	"database/sql"
	"github.com/giovanni-liboni/exercise-rest-api-shop/config"
	"github.com/giovanni-liboni/exercise-rest-api-shop/entities"
	"github.com/giovanni-liboni/exercise-rest-api-shop/repositories"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initOrderService(db *sql.DB) OrderService {
	// Create a new order repository
	orderRepo := repositories.InitOrderRepository(sqlx.NewDb(db, "mysql"))
	// Load the test environment
	configEnv := config.LoadConfig("../.test.env")
	// Create a new order service
	orderService := InitOrderService(orderRepo, configEnv.StripeAPIKey)

	return orderService
}

func TestOrderService_GetOrder(t *testing.T) {
	db, err := sql.Open("mysql_txdb", "testDataSource")
	if err != nil {
		t.Fatalf("mysql: failed to open a mysql connection, have you run 'make test'? err: %s", err)
	}
	defer db.Close()

	// Create a new order service
	orderService := initOrderService(db)

	// Get the order
	order, err := orderService.GetOrder(context.TODO(), 1)

	if err != nil {
		t.Fatalf("mysql: failed to get order, err: %s", err)
	}

	// Asserts
	assert.Equal(t, 1, order.ID)
}

func TestOrderService_PayOrder(t *testing.T) {
	db, err := sql.Open("mysql_txdb", "testDataSource")
	if err != nil {
		t.Fatalf("mysql: failed to open a mysql connection, have you run 'make test'? err: %s", err)
	}
	defer db.Close()

	// Create a new order service
	orderService := initOrderService(db)

	// Create the associated user
	user := entities.User{
		ID: 1,
		Email: "test@test.com",
		Firstname: "Test",
		Lastname: "User",
		Password: "test",
	}

	// Pay the order
	order, err := orderService.PayOrder(context.TODO(), 1, &user)
	if err != nil {
		t.Errorf("orderService.PayOrder() error = %v", err)
	}

	// Asserts
	assert.Equal(t, 1, order.ID)
	assert.Equal(t, "paid", order.Status)
}

func TestOrderService_GetOrders(t *testing.T) {
	db, err := sql.Open("mysql_txdb", "testDataSource")
	if err != nil {
		t.Fatalf("mysql: failed to open a mysql connection, have you run 'make test'? err: %s", err)
	}
	defer db.Close()

	// Create a new order service
	orderService := initOrderService(db)

	// Get the orders
	orders, err := orderService.GetOrders(context.TODO())

	if err != nil {
		t.Fatalf("mysql: failed to get orders, err: %s", err)
	}

	// Asserts
	assert.Equal(t, 1, len(orders))
}