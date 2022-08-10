package order_item

import (
	"context"

	shopRepo "github.com/Yu-Jack/shop-ddd-go/internal/adapter/repository/mysql/shop"
	"github.com/Yu-Jack/shop-ddd-go/internal/domain/shop"
)

type CreateOrderItemInput struct {
	Name       string
	Amount     int
	OrderID    string
	ConsumerID string
}

type Usecase interface {
	CreateOrderItem(ctx context.Context, input CreateOrderItemInput) (shop.OrderItem, error)
	GetOrderItems(ctx context.Context, orderId string) ([]shop.OrderItem, error)
}

type usecase struct {
	repo shopRepo.Repo
}

func NewUsecase(repo shopRepo.Repo) Usecase {
	return &usecase{
		repo: repo,
	}
}
