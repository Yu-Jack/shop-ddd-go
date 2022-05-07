package eventhandler

import (
	"github.com/Yu-Jack/shop-ddd-go/kit/logger"
)

func (eh *eventHandler) orderApproved() {
	go eh.eventBus.Subscribe("OrderApproved", func(orderId string) {
		ctx := logger.NewContext()
		eh.orderUsecase.ApproveOrder(ctx, orderId)
	})
}

func (eh *eventHandler) orderRejected() {
	go eh.eventBus.Subscribe("OrderRejected", func(orderId string) {
		ctx := logger.NewContext()
		eh.orderUsecase.RejectOrder(ctx, orderId)
	})
}
