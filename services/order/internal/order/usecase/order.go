package usecase

import (
	orderEntity "github.com/Yu-Jack/shop-ddd-go-order/internal/entity/order"
	"github.com/google/uuid"
)

func (u *usecase) CreateOrder(input CreateOrderInput) (*orderEntity.Order, error) {
	o := orderEntity.NewOrder()
	o.ID = uuid.NewString()
	o.UserID = input.UserID
	o.Name = input.Name
	o.State = "PENDING"
	o.Amount = 10 // fixed amount for demo
	// o.CreatedOrderEvent()
	u.repo.SaveOrder(o)
	// u.eventBus.Publish(o.DomainEvents)
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

func (u *usecase) FindOrderById(orderId string) (*orderEntity.Order, error) {
	o := u.repo.FindOrderByIds(orderId)
	return o, nil
}

func (u *usecase) FindOrderByConsumerId(consumerId string) (*orderEntity.Order, error) {
	o := u.repo.FindOrderByConsumerId(consumerId)
	return o, nil
}

func (u *usecase) GetAllOrders() ([]*orderEntity.Order, error) {
	os := u.repo.GetAllOrders()
	return os, nil
}
