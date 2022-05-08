package usecase

import (
	"context"
	"fmt"

	orderEntity "github.com/Yu-Jack/shop-ddd-go-order/internal/entity/order"
	"github.com/Yu-Jack/shop-ddd-go/kit/logger"
	"github.com/google/uuid"
)

func (u *usecase) CheckoutOrder(ctx context.Context, input CheckoutOrderInput, saga Saga) (orderEntity.Order, error) {
	o, err := u.repo.FindAvailableOrderByConsumerId(input.ConsumerID)
	if err != nil {
		logger.Log(ctx, "err", err)
		return o, err
	}

	amount, err := u.repo.FindTotalAmountByOrderId(o.ID)
	if err != nil {
		logger.Log(ctx, "err", err)
		return o, err
	}
	o.Amount = int(amount)
	o.State = "CHECKOUT_PENDING"
	u.repo.SaveOrder(o)
	o.CreatedOrderEvent()
	u.eventBus.Publish(o.DomainEvents)

	groupId := fmt.Sprintf("%s_%s", "order", o.ID)
	saga.CheckoutSaga(ctx, groupId, o.ID)

	return o, nil
}

func (u *usecase) CreateOrder(ctx context.Context, input CreateOrderInput) (orderEntity.Order, error) {
	o := orderEntity.NewOrder()
	o.ID = uuid.NewString()
	o.ConsumerID = input.ConsumerID
	o.Name = input.Name
	o.State = "PENDING"
	u.repo.CreateOrder(o)
	return *o, nil
}

func (u *usecase) ApproveOrder(ctx context.Context, orderId string) error {
	err := u.repo.UpdateOrderState(orderId, "APPROVED")
	return err
}

func (u *usecase) RejectOrder(ctx context.Context, orderId string) error {
	err := u.repo.UpdateOrderState(orderId, "REJECTED")
	return err
}

func (u *usecase) FindOrderById(ctx context.Context, orderId string) (orderEntity.Order, error) {

	o, err := u.repo.FindOrderById(orderId)
	if err != nil {
		logger.Log(ctx, "err", err)
		return o, err
	}
	return o, nil
}

func (u *usecase) FindAvailableOrderByConsumerId(ctx context.Context, consumerId string) (orderEntity.Order, error) {

	o, err := u.repo.FindAvailableOrderByConsumerId(consumerId)
	if err != nil {
		logger.Log(ctx, "err", err)
		return o, err
	}
	return o, nil
}

func (u *usecase) GetAllOrders(ctx context.Context) ([]orderEntity.Order, error) {

	os, err := u.repo.GetAllOrders()
	if err != nil {
		logger.Log(ctx, "err", err)
		return os, err
	}
	return os, nil
}
