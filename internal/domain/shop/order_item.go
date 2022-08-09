package order

import "github.com/Yu-Jack/shop-ddd-go/pkg/dddcore"

type OrderItem struct {
	ID           string
	OrderID      string
	Name         string
	Amount       int
	CreatedAt    int
	UpdatedAt    int
	DomainEvents []dddcore.Event
}
