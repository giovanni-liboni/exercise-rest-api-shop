package services

import (
	"context"
	"database/sql"
	"github.com/giovanni-liboni/exercise-rest-api-shop/repositories"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initItemService(db *sql.DB) ItemService {
	// Create a new order repository
	orderRepo := repositories.InitOrderRepository(sqlx.NewDb(db, "mysql"))

	// Create a new item repository
	itemRepo := repositories.InitItemRepository(sqlx.NewDb(db, "mysql"))

	// Create a new order service
	itemService := InitItemService(itemRepo, orderRepo)

	return itemService
}

func TestItemService_PurchaseItem(t *testing.T) {
	db, err := sql.Open("mysql_txdb", "testDataSource")
	if err != nil {
		t.Fatalf("mysql: failed to open a mysql connection, have you run 'make test'? err: %s", err)
	}
	defer db.Close()

	itemService := initItemService(db)

	// Purchase an item
	order, err := itemService.PurchaseItem(context.TODO(), 1, 1)

	if err != nil {
		t.Fatalf("mysql: failed to purchase an item, err: %s", err)
	}
	assert.Equal(t, int64(1), order.UserID)
}
