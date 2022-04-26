package usecase

import (
	"github.com/Yu-Jack/dddcore"
	"github.com/Yu-Jack/shop-ddd-go-consumer/entity"
	"github.com/Yu-Jack/shop-ddd-go-consumer/repository"
	"github.com/Yu-Jack/shop-ddd-go-consumer/request"
)

type Usecase interface {
	CheckOrder(orderId string, orderAmount int, consumerId string)
	CreateConsumer(req request.CreateConsumerRequest) entity.Consumer
	GetAllConsumers() []*entity.Consumer
}

type usecase struct {
	eventBus *dddcore.EventBus
	repo     repository.Repository
}

func New(eventBus *dddcore.EventBus, repo repository.Repository) Usecase {
	return &usecase{
		eventBus: eventBus,
		repo:     repo,
	}
}
