package eventhandler

import "fmt"

func (eh *eventHandler) orderApproved() {
	go eh.eventBus.Subscribe("OrderApproved", func(orderId string) {
		eh.orderUsecase.ApproveOrder(orderId)
	})
}

func (eh *eventHandler) orderRejected() {
	go eh.eventBus.Subscribe("OrderRejected", func(orderId string) {
		fmt.Println("OrderRejected")
		fmt.Println(orderId)
		eh.orderUsecase.RejectOrder(orderId)
	})
}
