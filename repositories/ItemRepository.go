package repositories

import (
	"context"
	"github.com/giovanni-liboni/exercise-rest-api-shop/entities"
	"github.com/jmoiron/sqlx"
)

type itemRepository struct {
	db *sqlx.DB
}

type ItemRepository interface {
	GetAllItems(ctx context.Context) ([]*entities.Item, error)
	GetItem(ctx context.Context, id int64) (*entities.Item, error)
	CreateItem(ctx context.Context, item *entities.Item) error
	UpdateItem(ctx context.Context, item *entities.Item) error
	DeleteItem(ctx context.Context, id int64) error
	GetItemsByOrderId(ctx context.Context, orderId int) ([]*entities.Item, error)
}

func InitItemRepository(db *sqlx.DB) ItemRepository {
	return &itemRepository{db}
}

func (i itemRepository) GetAllItems(ctx context.Context) ([]*entities.Item, error) {
	var items []*entities.Item
	err := i.db.SelectContext(ctx, &items, "CALL sp_GetItems")
	return items, err
}

func (i itemRepository) GetItem(ctx context.Context, id int64) (*entities.Item, error) {
	var item entities.Item
	err := i.db.GetContext(ctx, &item, "CALL sp_GetItem(?)", id)
	return &item, err
}

func (i itemRepository) CreateItem(ctx context.Context, item *entities.Item) error {
	_, err := i.db.NamedExecContext(ctx, "CALL sp_CreateItem(:name, :producer, :description, :price, :category)", item)
	return err
}

func (i itemRepository) UpdateItem(ctx context.Context, item *entities.Item) error {
	_, err := i.db.NamedExecContext(ctx, "CALL sp_UpdateItem(:id, :name, :producer, :description, :price, :category)", item)
	return err
}

func (i itemRepository) DeleteItem(ctx context.Context, id int64) error {
	_, err := i.db.ExecContext(ctx, "CALL sp_DeleteItem(?)", id)
	return err
}

func (i itemRepository) GetItemsByOrderId(ctx context.Context, orderId int) ([]*entities.Item, error) {
	var items []*entities.Item
	err := i.db.SelectContext(ctx, &items, "CALL sp_GetItemsByOrderId(?)", orderId)
	return items, err
}