package usecase

import (
	"context"

	orderEntity "github.com/Yu-Jack/shop-ddd-go-order/internal/entity/order"
	"github.com/Yu-Jack/shop-ddd-go/kit/logger"
	"github.com/google/uuid"
)

func (u *usecase) CreateOrderItem(ctx context.Context, input CreateOrderItemInput) (orderEntity.OrderItem, error) {
	log := logger.GetLogger(ctx)

	oi := &orderEntity.OrderItem{
		ID:      uuid.NewString(),
		Name:    input.Name,
		Amount:  input.Amount,
		OrderID: input.OrderID,
	}

	err := u.repo.CreateOrderItem(oi)
	if err != nil {
		log.Log("err", err)
		return orderEntity.OrderItem{}, err
	}

	return *oi, nil
}

func (u *usecase) GetOrderItems(ctx context.Context, orderId string) ([]orderEntity.OrderItem, error) {
	log := logger.GetLogger(ctx)

	ois, err := u.repo.GetAllOrderItemsByOrderId(orderId)
	if err != nil {
		log.Log("err", err)
		return ois, err
	}
	return ois, nil
}
