package usecase

import (
	"github.com/Yu-Jack/dddcore"
	orderEntity "github.com/Yu-Jack/shop-ddd-go-order/internal/entity/order"
)

type Usecase interface {
	CreateOrder(input CreateOrderInput) (*orderEntity.Order, error)
	ApproveOrder(orderId string) error
	RejectOrder(orderId string) error
	FindOrderById(orderId string) (*orderEntity.Order, error)
	FindOrderByConsumerId(consumerId string) (*orderEntity.Order, error)
	GetAllOrders() ([]*orderEntity.Order, error)
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