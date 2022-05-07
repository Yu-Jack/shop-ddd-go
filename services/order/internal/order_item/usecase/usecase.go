package usecase

import (
	"context"

	orderEntity "github.com/Yu-Jack/shop-ddd-go-order/internal/entity/order"
	"github.com/Yu-Jack/shop-ddd-go/kit/dddcore"
)

type Usecase interface {
	CreateOrderItem(ctx context.Context, input CreateOrderItemInput) (orderEntity.OrderItem, error)
	GetOrderItems(ctx context.Context, orderId string) ([]orderEntity.OrderItem, error)
}

type usecase struct {
	repo     Repository
	eventBus *dddcore.EventBus
}

func New(repo Repository, eventBus *dddcore.EventBus) Usecase {
	return &usecase{
		repo:     repo,
		eventBus: eventBus,
	}
}
