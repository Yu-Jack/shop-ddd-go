package shop

import (
	"context"

	repo "github.com/Yu-Jack/shop-ddd-go/internal/adapter/repository/mysql/shop"
	domain "github.com/Yu-Jack/shop-ddd-go/internal/domain/shop"
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
	repo repo.Repo
}

func NewOrderItem(repo repo.Repo) OrderItem {
	return &orderItem{
		repo: repo,
	}
}
