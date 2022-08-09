package order

import (
	"context"

	repo "github.com/Yu-Jack/shop-ddd-go/internal/adapter/repository/mysql/order"
	domain "github.com/Yu-Jack/shop-ddd-go/internal/domain/order"
)

type CreateOrderItemInput struct {
	Name       string
	Amount     int
	OrderID    string
	ConsumerID string
}

type OrderItem interface {
	CreateOrderItem(ctx context.Context, input CreateOrderItemInput) (domain.OrderItem, error)
	GetOrderItems(ctx context.Context, orderId string) ([]domain.OrderItem, error)
}

type orderItem struct {
	repo repo.Order
}

func NewOrderItem(repo repo.Order) OrderItem {
	return &orderItem{
		repo: repo,
	}
}
