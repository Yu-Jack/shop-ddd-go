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
	repo.db.Begin()
	defer repo.db.Commit()

	var c Consumer
	repo.db.Model(&Consumer{}).Where("id = ?", consumerId).Find(&c)

	if c.Amount < orderAmount {
		return errors.New("Consumer doesn't have enough money")
	}

	c.Amount -= orderAmount
	result := repo.db.Save(c)

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
