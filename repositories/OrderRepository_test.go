package repositories

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/giovanni-liboni/exercise-rest-api-shop/entities"
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

func TestOrderRepository_GetOrder(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Define the order rows to be returned by the database
	rows := sqlmock.NewRows([]string{"id", "user_id", "payment_method", "payment_id", "total_price", "status", "created_at", "updated_at"}).
		AddRow(1, 1, "stripe", "123456789", 100.00, "pending", time.Now(), time.Now())

	// Expect the query to be called with the correct SQL
	mock.ExpectQuery("CALL sp_GetOrder").WillReturnRows(rows)

	// Create a new OrderRepository with the database connection
	repo := InitOrderRepository(sqlx.NewDb(db, "mysql"))

	// Call the GetOrder method
	order, err := repo.GetOrder(context.TODO(), 1)

	// Check that the expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Assert that the returned order is not empty
	assert.NotEmpty(t, order)
	assert.Equal(t, int64(1), order.ID)
	assert.Equal(t, int64(1), order.UserID)
	assert.Equal(t, "stripe", order.PaymentMethod)
	assert.Equal(t, "123456789", order.PaymentID)
}

func TestOrderRepository_CreateOrder(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("CALL sp_CreateOrder").
		WithArgs(1, "paypal", "4382408943", 100.00, "pending").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Create a new OrderRepository with the database connection
	repo := InitOrderRepository(sqlx.NewDb(db, "mysql"))

	// Call the CreateOrder method
	order, err := repo.CreateOrder(context.TODO(), &entities.Order{
		UserID:        1,
		PaymentMethod: "paypal",
		PaymentID:     "4382408943",
		TotalPrice:    100.00,
		Status:        "pending",
	})

	if err != nil {
		t.Errorf("error was not expected while creating order: %s", err)
	}

	// Check that the expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.NotEmpty(t, order)
}

func TestOrderRepository_UpdateOrder(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("CALL sp_UpdateOrder").
		WithArgs(1, 1, "stripe", 100.0, "pending", "4382408943")

	// Create a new OrderRepository with the database connection
	repo := InitOrderRepository(sqlx.NewDb(db, "mysql"))

	// Call the UpdateOrder method
	err = repo.UpdateOrder(context.TODO(), &entities.Order{
		ID:            1,
		UserID:        1,
		PaymentMethod: "stripe",
		PaymentID:     "4382408943",
		TotalPrice:    100.00,
		Status:        "pending",
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
