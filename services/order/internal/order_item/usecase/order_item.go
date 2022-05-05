package usecase

import (
	orderEntity "github.com/Yu-Jack/shop-ddd-go-order/internal/entity/order"
	"github.com/google/uuid"
)

func (u *usecase) CreateOrderItem(input CreateOrderItemInput) (orderEntity.OrderItem, error) {
	oi := &orderEntity.OrderItem{
		ID:      uuid.NewString(),
		Name:    input.Name,
		Amount:  input.Amount,
		OrderID: input.OrderID,
	}
	u.repo.CreateOrderItem(oi)
	return *oi, nil
}

func (u *usecase) GetOrderItems(orderId string) ([]orderEntity.OrderItem, error) {
	ois, err := u.repo.GetAllOrderItemsByOrderId(orderId)
	if err != nil {
		return ois, err
	}
	return ois, nil
}
