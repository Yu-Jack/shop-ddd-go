package usecase

import (
	"context"

	orderEntity "github.com/Yu-Jack/shop-ddd-go-order/internal/entity/order"
	orderUc "github.com/Yu-Jack/shop-ddd-go-order/internal/order/usecase"
	"github.com/Yu-Jack/shop-ddd-go/kit/dddcore"
)

type Usecase interface {
	CreateOrderItem(ctx context.Context, input CreateOrderItemInput) (orderEntity.OrderItem, error)
	GetOrderItems(ctx context.Context, orderId string) ([]orderEntity.OrderItem, error)
}

type usecase struct {
	repo     orderUc.Repository
	eventBus *dddcore.EventBus
}

func New(repo orderUc.Repository, eventBus *dddcore.EventBus) Usecase {
	return &usecase{
		repo:     repo,
		eventBus: eventBus,
	}
}
