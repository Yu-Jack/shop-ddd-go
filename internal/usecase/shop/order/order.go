package order

import (
	"context"

	shopRepo "github.com/Yu-Jack/shop-ddd-go/internal/adapter/repository/mysql/shop"
	"github.com/Yu-Jack/shop-ddd-go/internal/domain/shop"
)

type CreateOrderInput struct {
	ConsumerID string
	Name       string
}

type CheckoutOrderInput struct {
	ConsumerID string
}

type Usecase interface {
	CheckoutOrder(ctx context.Context, input CheckoutOrderInput) (shop.Order, error)
	CreateOrder(ctx context.Context, input CreateOrderInput) (shop.Order, error)
	ApproveOrder(ctx context.Context, orderId string) error
	RejectOrder(ctx context.Context, orderId string) error
	FindOrderById(ctx context.Context, orderId string) (shop.Order, error)
	FindAvailableOrderByConsumerId(ctx context.Context, consumerId string) (shop.Order, error)
	GetAllOrders(ctx context.Context) ([]shop.Order, error)
}

type usecase struct {
	repo shopRepo.Repo
}

func NewUsecase(repo shopRepo.Repo) Usecase {
	return &usecase{
		repo: repo,
	}
}
