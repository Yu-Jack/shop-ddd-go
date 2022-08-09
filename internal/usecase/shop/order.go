package shop

import (
	"context"

	repo "github.com/Yu-Jack/shop-ddd-go/internal/adapter/repository/mysql/shop"
	domain "github.com/Yu-Jack/shop-ddd-go/internal/domain/shop"
)

type CreateOrderInput struct {
	ConsumerID string
	Name       string
}

type CheckoutOrderInput struct {
	ConsumerID string
}

type Order interface {
	CheckoutOrder(ctx context.Context, input CheckoutOrderInput) (domain.Order, error)
	CreateOrder(ctx context.Context, input CreateOrderInput) (domain.Order, error)
	ApproveOrder(ctx context.Context, orderId string) error
	RejectOrder(ctx context.Context, orderId string) error
	FindOrderById(ctx context.Context, orderId string) (domain.Order, error)
	FindAvailableOrderByConsumerId(ctx context.Context, consumerId string) (domain.Order, error)
	GetAllOrders(ctx context.Context) ([]domain.Order, error)
}

type order struct {
	repo repo.Order
}

func NewOrder(repo repo.Order) Order {
	return &order{
		repo: repo,
	}
}
