package usecase

import (
	orderEntity "github.com/Yu-Jack/shop-ddd-go-order/internal/entity/order"
	"github.com/google/uuid"
)

func (u *usecase) CheckoutOrder(input CheckoutOrderInput) (*orderEntity.Order, error) {
	o := u.repo.FindAvailableOrderByConsumerId(input.UserID)
	ois := u.repo.FindOrderItemsByOrderId(o.ID)
	totalAmount := 0
	for _, oi := range ois {
		totalAmount += oi.Amount
	}
	o.Amount = totalAmount
	o.CreatedOrderEvent()
	u.eventBus.Publish(o.DomainEvents)
	return o, nil
}

func (u *usecase) CreateOrder(input CreateOrderInput) (*orderEntity.Order, error) {
	o := orderEntity.NewOrder()
	o.ID = uuid.NewString()
	o.UserID = input.UserID
	o.Name = input.Name
	o.State = "PENDING"
	u.repo.SaveOrder(o)
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

func (u *usecase) FindAvailableOrderByConsumerId(consumerId string) (*orderEntity.Order, error) {
	o := u.repo.FindAvailableOrderByConsumerId(consumerId)
	return o, nil
}

func (u *usecase) GetAllOrders() ([]*orderEntity.Order, error) {
	os := u.repo.GetAllOrders()
	return os, nil
}
