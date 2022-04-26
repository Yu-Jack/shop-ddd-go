package repository

import "github.com/Yu-Jack/shop-ddd-go-consumer/entity"

var db = []*entity.Consumer{}

type repo struct{}

type Repository interface {
	DecreaseConsumerAmount(consumerId string, orderAmount int) (success bool)
	CreateConsumer(c entity.Consumer) error
	GetAllConsumers() ([]*entity.Consumer, error)
}

func New() Repository {
	return &repo{}
}
