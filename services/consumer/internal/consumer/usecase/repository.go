package usecase

import "github.com/Yu-Jack/shop-ddd-go-consumer/internal/entity"

type Repository interface {
	DecreaseConsumerAmount(consumerId string, orderAmount int) (success bool)
	CreateConsumer(c entity.Consumer) error
	GetAllConsumers() ([]*entity.Consumer, error)
}
