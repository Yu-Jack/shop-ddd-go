package usecase

import (
	"github.com/Yu-Jack/shop-ddd-go-order/entity"
	"github.com/google/uuid"
)

func (u *usecase) CreateOrder(input CreateOrderInput) (*entity.Order, error) {
	o := entity.NewOrder()
	o.ID = uuid.NewString()
	o.UserID = input.UserID
	o.Name = input.Name
	o.State = "PENDING"
	o.Amount = 10 // fixed amount for demo
	o.CreatedOrderEvent()
	u.repo.Save(o)
	u.eventBus.Publish(o.DomainEvents)
	return o, nil
}

func (u *usecase) ApproveOrder(orderId string) error {
	u.repo.UpdateOrderState(orderId, "APPROVED")
	return nil
}

func (u *usecase) RejectOrder(orderId string) error {
	u.repo.UpdateOrderState(orderId, "REJECTED")
	return nil
}

func (u *usecase) FindOrderById(orderId string) (*entity.Order, error) {
	o := u.repo.FindOrderByIds(orderId)
	return o, nil
}

func (u *usecase) GetAllOrders() ([]*entity.Order, error) {
	os := u.repo.GetAllOrders()
	return os, nil
}
