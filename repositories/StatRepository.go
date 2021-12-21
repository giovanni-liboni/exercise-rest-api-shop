package repositories

import (
	"context"
	"github.com/giovanni-liboni/exercise-rest-api-shop/entities"
	"github.com/jmoiron/sqlx"
)

type statRepository struct {
	db *sqlx.DB
}

type StatRepository interface {
	GetStatsLastMonth(ctx context.Context) (entities.Stat, error)
	GetStatsLastWeek(ctx context.Context) (entities.Stat, error)
	GetStatsLastDay(ctx context.Context) (entities.Stat, error)
	GetUsersWhoSpentMore(ctx context.Context) ([]entities.UserStat, error)
	GetMostOrderedItems(ctx context.Context) ([]entities.ItemStat, error)
	GetLeastOrderedItems(ctx context.Context) ([]entities.ItemStat, error)
	GetItemsNotOrdered(ctx context.Context) ([]entities.ItemStat, error)
	GetTotalNumberOfItems(ctx context.Context) (int, error)
	GetTotalNumberOfUsers(ctx context.Context) (int, error)
	GetTotalNumberOfOrders(ctx context.Context) (int, error)
}

func InitStatRepository(db *sqlx.DB) StatRepository {
	return &statRepository{db}
}

func (s statRepository) GetStatsLastMonth(ctx context.Context) (entities.Stat, error) {
	var stat entities.Stat
	err := s.db.GetContext(ctx, &stat, "CALL sp_GetStatLastMonth()")
	return  stat, err
}

func (s statRepository) GetStatsLastWeek(ctx context.Context) (entities.Stat, error) {
	var stat entities.Stat
	err := s.db.GetContext(ctx, &stat, "CALL sp_GetStatLastWeek()")
	return  stat, err
}

func (s statRepository) GetStatsLastDay(ctx context.Context) (entities.Stat, error) {
	var stat entities.Stat
	err := s.db.GetContext(ctx, &stat, "CALL sp_GetStatLastDay()")
	return  stat, err
}
func (s statRepository) GetUsersWhoSpentMore(ctx context.Context) ([]entities.UserStat, error) {
	var users []entities.UserStat
	err := s.db.SelectContext(ctx, &users, "CALL sp_GetUsersWhoSpentMore()")
	return  users, err
}

func (s statRepository) GetMostOrderedItems(ctx context.Context) ([]entities.ItemStat, error) {
	var items []entities.ItemStat
	err := s.db.SelectContext(ctx, &items, "CALL sp_GetMostOrderedItems()")
	return  items, err
}

func (s statRepository) GetLeastOrderedItems(ctx context.Context) ([]entities.ItemStat, error) {
	var items []entities.ItemStat
	err := s.db.SelectContext(ctx, &items, "CALL sp_GetLeastOrderedItems()")
	return  items, err
}

func (s statRepository) GetItemsNotOrdered(ctx context.Context) ([]entities.ItemStat, error) {
	var items []entities.ItemStat
	err := s.db.SelectContext(ctx, &items, "CALL sp_GetItemsNotOrdered()")
	return  items, err
}

func (s statRepository) GetTotalNumberOfItems(ctx context.Context) (int, error) {
	var total int
	err := s.db.GetContext(ctx, &total, "CALL sp_GetTotalNumberOfItems()")
	return  total, err
}

func (s statRepository) GetTotalNumberOfUsers(ctx context.Context) (int, error) {
	var total int
	err := s.db.GetContext(ctx, &total, "CALL sp_GetTotalNumberOfUsers()")
	return  total, err
}

func (s statRepository) GetTotalNumberOfOrders(ctx context.Context) (int, error) {
	var total int
	err := s.db.GetContext(ctx, &total, "CALL sp_GetTotalNumberOfOrders()")
	return  total, err
}