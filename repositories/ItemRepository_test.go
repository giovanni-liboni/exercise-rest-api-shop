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

func TestItemRepository_GetAllItems(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Define the rows to be returned by the database
	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "producer", "category", "created_at", "updated_at"}).
		AddRow(1, "Item 1", "Description 1", 1.0, "Producer 1", "Category 1", time.Now(), time.Now()).
		AddRow(2, "Item 2", "Description 2", 2.0, "Producer 2", "Category 2", time.Now(), time.Now())

	// Expect the query to be called with the correct SQL
	mock.ExpectQuery("CALL sp_GetItems").WillReturnRows(rows)

	// Create a new ItemRepository with the database connection
	repo := InitItemRepository(sqlx.NewDb(db, "mysql"))

	// Call the GetAllItems method
	items, err := repo.GetAllItems(context.TODO())
	if err != nil {
		t.Errorf("error was not expected while getting items: %s", err)
	}

	// Check that the expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Assert that the returned items are not empty
	assert.NotEmpty(t, items)
}

func TestItemRepository_GetItem(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Define the rows to be returned by the database
	timestamp := time.Now()
	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "producer", "category", "created_at", "updated_at"}).
		AddRow(1, "Item 1", "Description 1", 1.0, "Producer 1", "Category 1", timestamp, timestamp)

	// Expect the database to be called with the correct query
	mock.ExpectQuery("CALL sp_GetItem").WithArgs(1).WillReturnRows(rows)

	// Create a new ItemRepository with the database connection
	repo := InitItemRepository(sqlx.NewDb(db, "mysql"))

	// Call the repository method
	item, err := repo.GetItem(context.TODO(), 1)
	if err != nil {
		t.Errorf("error was not expected while getting item: %s", err)
	}

	// Check that the expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Check if the returned item has the correct values
	assert.Equal(t, int64(1), item.ID)
	assert.Equal(t, "Item 1", item.Name)
	assert.Equal(t, "Description 1", item.Description)
	assert.Equal(t, float32(1.0), item.Price)
	assert.Equal(t, "Producer 1", item.Producer)
	assert.Equal(t, "Category 1", item.Category)
	assert.Equal(t, timestamp, item.CreatedAt)
	assert.Equal(t, timestamp, item.UpdatedAt)
}

func TestItemRepository_CreateItem(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("CALL sp_CreateItem").
		WithArgs("Sample item", "Sample producer", "Sample description", 1.0, "Sample category").
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := InitItemRepository(sqlx.NewDb(db, "mysql"))

	item := &entities.Item{
		Name:        "Sample item",
		Description: "Sample description",
		Price:       1.0,
		Producer:    "Sample producer",
		Category:    "Sample category",
	}

	err = repo.CreateItem(context.TODO(), item)
	if err != nil {
		t.Errorf("error was not expected while creating item: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, "Sample item", item.Name)
	assert.Equal(t, "Sample description", item.Description)
	assert.Equal(t, float32(1.0), item.Price)
	assert.Equal(t, "Sample producer", item.Producer)
	assert.Equal(t, "Sample category", item.Category)

}

func TestItemRepository_UpdateItem(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("CALL sp_UpdateItem").
		WithArgs(1, "Sample item", "Sample producer", "Sample description", 1.0, "Sample category").
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := InitItemRepository(sqlx.NewDb(db, "mysql"))

	item := &entities.Item{
		ID:          1,
		Name:        "Sample item",
		Description: "Sample description",
		Price:       1.0,
		Producer:    "Sample producer",
		Category:    "Sample category",
	}

	err = repo.UpdateItem(context.TODO(), item)

	if err != nil {
		t.Errorf("error was not expected while updating item: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	assert.Equal(t, int64(1), item.ID)
	assert.Equal(t, "Sample item", item.Name)
	assert.Equal(t, "Sample description", item.Description)
	assert.Equal(t, float32(1.0), item.Price)
	assert.Equal(t, "Sample producer", item.Producer)
	assert.Equal(t, "Sample category", item.Category)

}

func TestItemRepository_DeleteItem(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("CALL sp_DeleteItem").
		WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := InitItemRepository(sqlx.NewDb(db, "mysql"))

	err = repo.DeleteItem(context.TODO(), 1)

	if err != nil {
		t.Errorf("error was not expected while deleting item: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestItemRepository_GetItemsByOrderId(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("CALL sp_GetItemsByOrderId").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "price", "producer", "category", "created_at", "updated_at"}).
			AddRow(1, "Item 1", "Description 1", 1.0, "Producer 1", "Category 1", time.Now(), time.Now()).
			AddRow(2, "Item 2", "Description 2", 2.0, "Producer 2", "Category 2", time.Now(), time.Now()))

	repo := InitItemRepository(sqlx.NewDb(db, "mysql"))

	items, err := repo.GetItemsByOrderId(context.TODO(), 1)

	if err != nil {
		t.Errorf("error was not expected while getting items by order id: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 2, len(items))
}
