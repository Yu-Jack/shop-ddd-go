package usecase

import (
	"context"

	orderEntity "github.com/Yu-Jack/shop-ddd-go-order/internal/entity/order"
	"github.com/Yu-Jack/shop-ddd-go/kit/dddcore"
)

type Usecase interface {
	CheckoutOrder(ctx context.Context, input CheckoutOrderInput, saga Saga) (orderEntity.Order, error)
	CreateOrder(ctx context.Context, input CreateOrderInput) (orderEntity.Order, error)
	ApproveOrder(ctx context.Context, orderId string) error
	RejectOrder(ctx context.Context, orderId string) error
	FindOrderById(ctx context.Context, orderId string) (orderEntity.Order, error)
	FindAvailableOrderByConsumerId(ctx context.Context, consumerId string) (orderEntity.Order, error)
	GetAllOrders(ctx context.Context) ([]orderEntity.Order, error)
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
