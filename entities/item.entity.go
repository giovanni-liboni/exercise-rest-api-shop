package entities

import "time"

type Item struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	Producer    string    `json:"producer"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type ItemStat struct {
	Item
	TotalOrders int64 `json:"total_orders" db:"total_orders"`
}
