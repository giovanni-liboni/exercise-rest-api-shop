package entities

type Stat struct {
	TotalAmount float64 `json:"totalAmount" db:"total_amount"`
	TotalOrders int64   `json:"totalOrders" db:"total_orders"`
	TotalUsers  int64   `json:"totalUsers" db:"total_users"`
}
