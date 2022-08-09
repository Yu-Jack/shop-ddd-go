package order

import (
	"context"

	domain "github.com/Yu-Jack/shop-ddd-go/internal/domain/order"
	"github.com/Yu-Jack/shop-ddd-go/pkg/logger"
	"github.com/google/uuid"
)

func (usecase *order) CheckoutOrder(ctx context.Context, input CheckoutOrderInput) (domain.Order, error) {
	o, err := usecase.repo.FindAvailableOrderByConsumerId(input.ConsumerID)
	if err != nil {
		logger.Log(ctx, "err", err)
		return o, err
	}

	amount, err := usecase.repo.FindTotalAmountByOrderId(o.ID)
	if err != nil {
		logger.Log(ctx, "err", err)
		return o, err
	}
	o.Amount = int(amount)
	o.State = "CHECKOUT_PENDING"
	usecase.repo.SaveOrder(o)
	// o.CreatedOrderEvent()
	// usecase.eventBus.Publish(o.DomainEvents)

	// groupId := fmt.Sprintf("%s_%s", "order", o.ID)
	// saga.CheckoutSaga(ctx, groupId, o.ID)

	return o, nil
}

func (usecase *order) CreateOrder(ctx context.Context, input CreateOrderInput) (domain.Order, error) {
	o := domain.Order{}
	o.ID = uuid.NewString()
	o.ConsumerID = input.ConsumerID
	o.Name = input.Name
	o.State = "PENDING"
	usecase.repo.CreateOrder(&o)
	return o, nil
}

func (usecase *order) ApproveOrder(ctx context.Context, orderId string) error {
	err := usecase.repo.UpdateOrderState(orderId, "APPROVED")
	return err
}

func (usecase *order) RejectOrder(ctx context.Context, orderId string) error {
	err := usecase.repo.UpdateOrderState(orderId, "REJECTED")
	return err
}

func (usecase *order) FindOrderById(ctx context.Context, orderId string) (domain.Order, error) {

	o, err := usecase.repo.FindOrderById(orderId)
	if err != nil {
		logger.Log(ctx, "err", err)
		return o, err
	}
	return o, nil
}

func (usecase *order) FindAvailableOrderByConsumerId(ctx context.Context, consumerId string) (domain.Order, error) {

	o, err := usecase.repo.FindAvailableOrderByConsumerId(consumerId)
	if err != nil {
		logger.Log(ctx, "err", err)
		return o, err
	}
	return o, nil
}

func (usecase *order) GetAllOrders(ctx context.Context) ([]domain.Order, error) {

	os, err := usecase.repo.GetAllOrders()
	if err != nil {
		logger.Log(ctx, "err", err)
		return os, err
	}
	return os, nil
}
