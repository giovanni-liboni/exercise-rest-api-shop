package entities

import (
	"encoding/json"
	"strconv"
	"time"
)

type Item struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       StringFloat32   `json:"price"`
	Producer    string    `json:"producer"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type ItemStat struct {
	Item
	TotalOrders int64 `json:"total_orders" db:"total_orders"`
}

// StringInt create a type alias for type int
type StringFloat32 float32

// UnmarshalJSON create a custom unmarshal for the StringInt
/// this helps us check the type of our value before unmarshalling it

func (st *StringFloat32) UnmarshalJSON(b []byte) error {
	//convert the bytes into an interface
	//this will help us check the type of our value
	//if it is a string that can be converted into an int we convert it
	///otherwise we return an error
	var item interface{}
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	float, err := strconv.ParseFloat(item.(string), 32)
	if err != nil {
		return err
	}
	*st = StringFloat32(float)

	return nil
}