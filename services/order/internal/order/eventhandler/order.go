package eventhandler

import "context"

func (eh *eventHandler) orderApproved() {
	go eh.eventBus.Subscribe("OrderApproved", func(orderId string) {
		eh.orderUsecase.ApproveOrder(context.Background(), orderId)
	})
}

func (eh *eventHandler) orderRejected() {
	go eh.eventBus.Subscribe("OrderRejected", func(orderId string) {
		eh.orderUsecase.RejectOrder(context.Background(), orderId)
	})
}
