package entities

import "time"

type UserStat struct {
	ID        int64       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Firstname string    `json:"firstname" db:"firstname"`
	Lastname  string    `json:"lastname" db:"lastname"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Role      string    `json:"role" db:"role"`
	Orders    []Order   `json:"orders"`
	TotalSpent float64   `json:"total_spent" db:"total_spent"`
}