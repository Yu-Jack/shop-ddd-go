package order

import (
	"github.com/Yu-Jack/dddcore"
	"github.com/Yu-Jack/shop-ddd-go-order/entity"
)

type Usecase interface {
	CreateOrder(input CreateOrderInput) (*entity.Order, error)
	ApproveOrder(orderId string) error
	RejectOrder(orderId string) error
	FindOrderById(orderId string) (*entity.Order, error)
	GetAllOrders() ([]*entity.Order, error)
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
