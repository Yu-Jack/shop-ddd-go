package order_item

import (
	"context"

	"github.com/Yu-Jack/shop-ddd-go/internal/domain/shop"
	"github.com/Yu-Jack/shop-ddd-go/pkg/logger"
	"github.com/google/uuid"
)

func (usecase *usecase) CreateOrderItem(ctx context.Context, input CreateOrderItemInput) (shop.OrderItem, error) {
	oi := &shop.OrderItem{
		ID:      uuid.NewString(),
		Name:    input.Name,
		Amount:  input.Amount,
		OrderID: input.OrderID,
	}

	err := usecase.repo.CreateOrderItem(oi, input.ConsumerID)
	if err != nil {
		logger.Log(ctx, "err", err)
		return shop.OrderItem{}, err
	}

	return *oi, nil
}

func (usecase *usecase) GetOrderItems(ctx context.Context, orderId string) ([]shop.OrderItem, error) {

	// ois, err := u.repo.GetAllOrderItemsByOrderId(orderId)
	// if err != nil {
	// 	logger.Log(ctx, "err", err)
	// 	return ois, err
	// }
	return []shop.OrderItem{}, nil
}
