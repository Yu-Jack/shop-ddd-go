package shop

import "github.com/Yu-Jack/shop-ddd-go/pkg/dddcore"

type Order struct {
	ID           string
	ConsumerID   string
	Name         string
	State        string
	Amount       int
	CreatedAt    int
	UpdatedAt    int
	DomainEvents []dddcore.Event
}
