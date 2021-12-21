package services

import (
	"context"
	"github.com/giovanni-liboni/exercise-rest-api-shop/entities"
	"github.com/giovanni-liboni/exercise-rest-api-shop/repositories"
	log "github.com/sirupsen/logrus"
)

type itemService struct {
	itemRepository  repositories.ItemRepository
	orderRepository repositories.OrderRepository
}

type ItemService interface {
	GetAllItems(ctx context.Context) ([]*entities.Item, error)
	GetItem(ctx context.Context, id int64) (*entities.Item, error)
	CreateItem(ctx context.Context, item *entities.Item) error
	UpdateItem(ctx context.Context, item *entities.Item) error
	DeleteItem(ctx context.Context, id int64) error
	PurchaseItem(ctx context.Context, id int64, userId int64) (*entities.Order, error)
}

func InitItemService(itemRepository repositories.ItemRepository, orderRepository repositories.OrderRepository) ItemService {
	return &itemService{
		itemRepository:  itemRepository,
		orderRepository: orderRepository,
	}
}

func (i itemService) GetAllItems(ctx context.Context) ([]*entities.Item, error) {
	return i.itemRepository.GetAllItems(ctx)
}

func (i itemService) GetItem(ctx context.Context, id int64) (*entities.Item, error) {
	return i.itemRepository.GetItem(ctx, id)
}

func (i itemService) CreateItem(ctx context.Context, item *entities.Item) error {
	return i.itemRepository.CreateItem(ctx, item)
}

func (i itemService) UpdateItem(ctx context.Context, item *entities.Item) error {
	return i.itemRepository.UpdateItem(ctx, item)
}

func (i itemService) DeleteItem(ctx context.Context, id int64) error {
	return i.itemRepository.DeleteItem(ctx, id)
}

// PurchaseItem is a method that allows to purchase an item by a user
// It creates a new order and adds the item to the order
func (i itemService) PurchaseItem(ctx context.Context, id int64, userId int64) (*entities.Order, error) {
	// Get the item
	item, err := i.itemRepository.GetItem(ctx, id)
	if err != nil {
		return nil, err
	}

	// Create a new order
	order := entities.Order{
		UserID: userId,
		Items: []*entities.Item{
			item,
		},
		Status:        entities.OrderStatusCreated,
		TotalPrice:    item.Price,
		PaymentID:     "",
		PaymentMethod: "stripe",
	}

	// Create the order
	_, err = i.orderRepository.CreateOrder(ctx, &order)
	if err != nil {
		log.Errorf("Error creating order: %v", err)
		return nil, err
	}

	return &order, nil
}
