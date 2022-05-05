package order

import (
	"github.com/Yu-Jack/dddcore"
)

type OrderItem struct {
	ID           string          `json:"id" gorm:"primarykey"`
	OrderID      string          `json:"order_id"`
	Name         string          `json:"name"`
	Amount       int             `json:"amount"`
	CreatedAt    int             `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt    int             `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DomainEvents []dddcore.Event `json:"-" gorm:"-"`
}
