package entity

import (
	"encoding/json"
)

type Order struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	State  string `json:"state"`
	Amount int    `json:"amount"`
}

func NewOrder() Order {
	return Order{}
}

func (o Order) ToJsonString() string {
	data, _ := json.Marshal(o)
	return string(data)
}
