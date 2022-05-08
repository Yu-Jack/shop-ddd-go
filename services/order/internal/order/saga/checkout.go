package saga

import (
	"context"
	"fmt"

	orderUc "github.com/Yu-Jack/shop-ddd-go-order/internal/order/usecase"
	"github.com/Yu-Jack/shop-ddd-go/kit/dddcore"
	"github.com/Yu-Jack/shop-ddd-go/kit/logger"
)

type saga struct {
	eventBus *dddcore.EventBus
	orderUc  orderUc.Usecase
}

func NewSaga(eventBus *dddcore.EventBus, orderUc orderUc.Usecase) *saga {
	return &saga{
		eventBus: eventBus,
		orderUc:  orderUc,
	}
}

func (saga *saga) CheckoutSaga(ctx context.Context, groupId string, eventUUID string) {
	s := dddcore.NewSaga(ctx, groupId, saga.eventBus)

	s.AddStep(dddcore.SagaData{
		Name: "Validate if consumer is available to checkout",
		Invoke: func(ctx context.Context, orderId string) {
			saga.orderUc.ApproveOrder(ctx, orderId)
		},
		InvokeKey: fmt.Sprintf("%s-%s", "OrderApproved", eventUUID),
		Compensation: func(ctx context.Context, orderId string) {
			saga.orderUc.RejectOrder(ctx, orderId)
		},
		CompensationKey: fmt.Sprintf("%s-%s", "OrderRejected", eventUUID),
	})

	go func() {
		err := s.Execute()
		if err != nil {
			logger.Log(ctx, "err", err)
		}
	}()
}
