package usecase

import (
	orderEntity "github.com/Yu-Jack/shop-ddd-go-order/internal/entity/order"
	"github.com/Yu-Jack/shop-ddd-go/kit/dddcore"
)

type Usecase interface {
	CheckoutOrder(input CheckoutOrderInput) (orderEntity.Order, error)
	CreateOrder(input CreateOrderInput) (orderEntity.Order, error)
	ApproveOrder(orderId string) error
	RejectOrder(orderId string) error
	FindOrderById(orderId string) (orderEntity.Order, error)
	FindAvailableOrderByConsumerId(consumerId string) (orderEntity.Order, error)
	GetAllOrders() ([]orderEntity.Order, error)
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
