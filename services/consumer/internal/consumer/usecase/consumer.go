package usecase

import (
	"github.com/Yu-Jack/shop-ddd-go-consumer/internal/entity"
	"github.com/google/uuid"
)

func (u *usecase) CreateConsumer(input CreateConsumerInput) entity.Consumer {
	c := entity.Consumer{
		ID:        uuid.NewString(),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Amount:    input.Amount,
	}
	u.repo.CreateConsumer(c)
	return c
}

func (u *usecase) GetAllConsumers() []*entity.Consumer {
	cs, _ := u.repo.GetAllConsumers()
	return cs
}
