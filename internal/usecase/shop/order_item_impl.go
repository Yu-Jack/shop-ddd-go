package order

import (
	"context"

	domain "github.com/Yu-Jack/shop-ddd-go/internal/domain/shop"
	"github.com/Yu-Jack/shop-ddd-go/pkg/logger"
	"github.com/google/uuid"
)

func (usecase *orderItem) CreateOrderItem(ctx context.Context, input CreateOrderItemInput) (domain.OrderItem, error) {
	oi := &domain.OrderItem{
		ID:      uuid.NewString(),
		Name:    input.Name,
		Amount:  input.Amount,
		OrderID: input.OrderID,
	}

	err := usecase.repo.CreateOrderItem(oi, input.ConsumerID)
	if err != nil {
		logger.Log(ctx, "err", err)
		return domain.OrderItem{}, err
	}

	return *oi, nil
}

func (usecase *orderItem) GetOrderItems(ctx context.Context, orderId string) ([]domain.OrderItem, error) {

	// ois, err := u.repo.GetAllOrderItemsByOrderId(orderId)
	// if err != nil {
	// 	logger.Log(ctx, "err", err)
	// 	return ois, err
	// }
	return []domain.OrderItem{}, nil
}
