package usecase

import (
	orderEntity "github.com/Yu-Jack/shop-ddd-go-order/internal/entity/order"
	"github.com/google/uuid"
)

func (u *usecase) CreateOrderItem(input CreateOrderItemInput) (orderEntity.OrderItem, error) {
	oi := orderEntity.OrderItem{
		ID:      uuid.NewString(),
		Name:    input.Name,
		Amount:  input.Amount,
		OrderID: input.OrderID,
	}
	u.repo.SaveOrderItem(oi)
	return oi, nil
}
