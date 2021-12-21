package services

import (
	"context"
	"github.com/giovanni-liboni/exercise-rest-api-shop/entities"
	"github.com/giovanni-liboni/exercise-rest-api-shop/repositories"
	log "github.com/sirupsen/logrus"
)

type statService struct {
	userRepository repositories.UserRepository
	orderRepository repositories.OrderRepository
	itemRepository repositories.ItemRepository
	statRepository repositories.StatRepository
}

type StatService interface {
	GetItemsNotOrdered(ctx context.Context) ([]entities.ItemStat, error)
	GetUsersWhoSpendMore(ctx context.Context) ([]entities.UserStat, error)
	GetMostOrderedItems(ctx context.Context) ([]entities.ItemStat, error)
	GetStats(ctx context.Context) ([]entities.Stat, error)
	GetLeastOrderedItems(ctx context.Context) ([]entities.ItemStat, error)
	GetTotalNumberOfItems(ctx context.Context) (int, error)
	GetTotalNumberOfUsers(ctx context.Context) (int, error)
	GetTotalNumberOfOrders(ctx context.Context) (int, error)
}

func InitStatService(userRepository repositories.UserRepository, orderRepository repositories.OrderRepository, itemRepository repositories.ItemRepository, statRepository repositories.StatRepository) StatService {
	return &statService{
		userRepository: userRepository,
		orderRepository: orderRepository,
		itemRepository: itemRepository,
		statRepository: statRepository,
	}
}

func (s statService) GetLeastOrderedItems(ctx context.Context) ([]entities.ItemStat, error) {
	return s.statRepository.GetLeastOrderedItems(ctx)
}

func (s statService) GetItemsNotOrdered(ctx context.Context) ([]entities.ItemStat, error) {
	return s.statRepository.GetItemsNotOrdered(ctx)
}

func (s statService) GetUsersWhoSpendMore(ctx context.Context) ([]entities.UserStat, error) {
	return s.statRepository.GetUsersWhoSpentMore(ctx)
}

func (s statService) GetMostOrderedItems(ctx context.Context) ([]entities.ItemStat, error) {
	return s.statRepository.GetMostOrderedItems(ctx)
}

// GetStats returns the stats of the website
// In position 0 there are the stats for the month,
// in position 1 there are the stats for the week,
// in position 2 there are the stats for the last day.
func (s statService) GetStats(ctx context.Context) ([]entities.Stat, error) {
	// Get stat for last month
	lastMonthStat, err := s.statRepository.GetStatsLastMonth(ctx)
	if err != nil {
		log.Errorln("Error getting stats for last month: ", err)
		return nil, err
	}

	// Get stat for last week
	lastWeekStat, err := s.statRepository.GetStatsLastWeek(ctx)
	if err != nil {
		log.Errorln("Error getting stats for last week: ", err)
		return nil, err
	}

	// Get stat for last day
	lastDayStat, err := s.statRepository.GetStatsLastDay(ctx)
	if err != nil {
		log.Errorln("Error getting stats for last day: ", err)
		return nil, err
	}

	// Concatenate all stats
	stats := []entities.Stat{}
	stats = append(stats, lastMonthStat)
	stats = append(stats, lastWeekStat)
	stats = append(stats, lastDayStat)

	return stats, nil
}

func (s statService) GetTotalNumberOfItems(ctx context.Context) (int, error) {
	return s.statRepository.GetTotalNumberOfItems(ctx)
}

func (s statService) GetTotalNumberOfUsers(ctx context.Context) (int, error) {
	return s.statRepository.GetTotalNumberOfUsers(ctx)
}

func (s statService) GetTotalNumberOfOrders(ctx context.Context) (int, error) {
	return s.statRepository.GetTotalNumberOfOrders(ctx)
}