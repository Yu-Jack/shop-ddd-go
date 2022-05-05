package usecase

import (
	"github.com/Yu-Jack/shop-ddd-go-consumer/internal/entity"
	"github.com/Yu-Jack/shop-ddd-go/kit/dddcore"
)

type Usecase interface {
	CheckOrder(orderId string, orderAmount int, consumerId string)
	CreateConsumer(input CreateConsumerInput) entity.Consumer
	GetAllConsumers() []*entity.Consumer
}

type usecase struct {
	eventBus *dddcore.EventBus
	repo     Repository
}

func New(eventBus *dddcore.EventBus, repo Repository) Usecase {
	return &usecase{
		eventBus: eventBus,
		repo:     repo,
	}
}
