package order

import "github.com/Yu-Jack/dddcore"

type OrderItem struct {
	ID           string          `json:"id"`
	OrderID      string          `json:"order_id"`
	Name         string          `json:"name"`
	Amount       int             `json:"amount"`
	DomainEvents []dddcore.Event `json:"-"`
}
