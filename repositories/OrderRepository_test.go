package repositories

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestOrderRepository_GetAllOrders(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	timestamp := time.Now()
	// Define the order rows to be returned by the database
	rows := sqlmock.NewRows([]string{"id", "user_id", "payment_method", "payment_id", "total_price", "status", "created_at", "updated_at"}).
		AddRow(1, 1, "paypal", 1, 100.00, "pending", timestamp, timestamp).
		AddRow(2, 1, "paypal", 1, 100.00, "pending", timestamp, timestamp)


	// Expect the query to be called with the correct SQL
	mock.ExpectQuery("CALL sp_GetOrders").WillReturnRows(rows)

	// Create a new ItemRepository with the database connection
	repo := InitOrderRepository(sqlx.NewDb(db, "mysql"))

	// Call the GetAllItems method
	items, err := repo.GetAllOrders(context.TODO())
	if err != nil {
		t.Errorf("error was not expected while getting users: %s", err)
	}

	// Check that the expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Assert that the returned items are not empty
	assert.NotEmpty(t, items)
}