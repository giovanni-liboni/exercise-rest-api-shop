package services

import (
	"context"
	"github.com/giovanni-liboni/exercise-rest-api-shop/entities"
	"github.com/giovanni-liboni/exercise-rest-api-shop/repositories"
)

type itemService struct {
	itemRepository repositories.ItemRepository
}

type ItemService interface {
	GetAllItems(ctx context.Context) ([]*entities.Item, error)
	GetItem(ctx context.Context, id int) (*entities.Item, error)
	CreateItem(ctx context.Context, item *entities.Item) error
	UpdateItem(ctx context.Context, item *entities.Item) error
	DeleteItem(ctx context.Context, id int) error
	PurchaseItem(ctx context.Context, id int, userId int) error
}

func InitItemService(itemRepository repositories.ItemRepository) ItemService {
	return &itemService{
		itemRepository: itemRepository,
	}
}

func (i itemService) GetAllItems(ctx context.Context) ([]*entities.Item, error) {
	return i.itemRepository.GetAllItems(ctx)
}

func (i itemService) GetItem(ctx context.Context, id int) (*entities.Item, error) {
	return i.itemRepository.GetItem(ctx, id)
}

func (i itemService) CreateItem(ctx context.Context, item *entities.Item) error {
	return i.itemRepository.CreateItem(ctx, item)
}

func (i itemService) UpdateItem(ctx context.Context, item *entities.Item) error {
	return i.itemRepository.UpdateItem(ctx, item)
}

func (i itemService) DeleteItem(ctx context.Context, id int) error {
	return i.itemRepository.DeleteItem(ctx, id)
}

func (i itemService) PurchaseItem(ctx context.Context, id int, userId int) error {
	panic("implement me")
}
