package repository

import (
	"github.com/Yu-Jack/shop-ddd-go-consumer/entity"
)

func (repo *repo) DecreaseConsumerAmount(consumerId string, orderAmount int) bool {
	for _, c := range db {
		if c.ID == consumerId && c.Amount > orderAmount {
			c.Amount = c.Amount - orderAmount
			return true
		}
	}
	return false
}

func (repo *repo) CreateConsumer(c entity.Consumer) error {
	db = append(db, &c)
	return nil
}

func (repo *repo) GetAllConsumers() ([]*entity.Consumer, error) {
	return db, nil
}
