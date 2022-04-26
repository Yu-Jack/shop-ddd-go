package usecase

import (
	"github.com/Yu-Jack/shop-ddd-go-consumer/entity"
	"github.com/Yu-Jack/shop-ddd-go-consumer/request"
	"github.com/google/uuid"
)

func (u *usecase) CreateConsumer(req request.CreateConsumerRequest) entity.Consumer {
	c := entity.Consumer{
		ID:        uuid.NewString(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Amount:    req.Amount,
	}
	u.repo.CreateConsumer(c)
	return c
}

func (u *usecase) GetAllConsumers() []*entity.Consumer {
	cs, _ := u.repo.GetAllConsumers()
	return cs
}
