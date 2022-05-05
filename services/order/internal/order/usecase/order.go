package usecase

import (
	orderEntity "github.com/Yu-Jack/shop-ddd-go-order/internal/entity/order"
	"github.com/google/uuid"
)

func (u *usecase) CheckoutOrder(input CheckoutOrderInput) (orderEntity.Order, error) {
	o, err := u.repo.FindAvailableOrderByConsumerId(input.ConsumerID)
	if err != nil {
		return o, err
	}

	amount, err := u.repo.FindTotalAmountByOrderId(o.ID)
	if err != nil {
		return o, err
	}
	o.Amount = int(amount)
	o.State = "CHECKOUT_PENDING"
	u.repo.SaveOrder(o)
	o.CreatedOrderEvent()
	u.eventBus.Publish(o.DomainEvents)
	return o, nil
}

func (u *usecase) CreateOrder(input CreateOrderInput) (orderEntity.Order, error) {
	o := orderEntity.NewOrder()
	o.ID = uuid.NewString()
	o.ConsumerID = input.ConsumerID
	o.Name = input.Name
	o.State = "PENDING"
	u.repo.CreateOrder(o)
	return *o, nil
}

func (u *usecase) ApproveOrder(orderId string) error {
	err := u.repo.UpdateOrderState(orderId, "APPROVED")
	return err
}

func (u *usecase) RejectOrder(orderId string) error {
	err := u.repo.UpdateOrderState(orderId, "REJECTED")
	return err
}

func (u *usecase) FindOrderById(orderId string) (orderEntity.Order, error) {
	o, err := u.repo.FindOrderById(orderId)
	if err != nil {
		return o, err
	}
	return o, nil
}

func (u *usecase) FindAvailableOrderByConsumerId(consumerId string) (orderEntity.Order, error) {
	o, err := u.repo.FindAvailableOrderByConsumerId(consumerId)
	if err != nil {
		return o, err
	}
	return o, nil
}

func (u *usecase) GetAllOrders() ([]orderEntity.Order, error) {
	os, err := u.repo.GetAllOrders()
	if err != nil {
		return os, err
	}
	return os, nil
}
