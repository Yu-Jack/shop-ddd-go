package entity

import (
	"errors"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *repo {
	return &repo{
		db: db,
	}
}

func (repo *repo) DecreaseConsumerAmount(consumerId string, orderAmount int) error {
	tx := repo.db.Begin()

	var c Consumer
	tx.Model(&Consumer{}).Where("id = ?", consumerId).Find(&c)

	if c.Amount < orderAmount {
		return errors.New("Consumer doesn't have enough money")
	}
	c.Amount -= orderAmount
	tx.Save(c)

	result := tx.Commit()

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *repo) CreateConsumer(c *Consumer) error {
	result := repo.db.Create(c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *repo) GetAllConsumers() (cus []Consumer, err error) {
	result := repo.db.Model(&Consumer{}).Find(&cus)

	if result.Error != nil {
		return cus, result.Error
	}

	return cus, nil
}
