package eventhandler

import (
	"github.com/Yu-Jack/dddcore"
	ucOrder "github.com/Yu-Jack/shop-ddd-go-order/usecase/order"
)

type eventHandler struct {
	orderUsecase ucOrder.Usecase
	eventBus     *dddcore.EventBus
}

func New(orderUsecase ucOrder.Usecase, eventBus *dddcore.EventBus) *eventHandler {
	return &eventHandler{
		orderUsecase: orderUsecase,
		eventBus:     eventBus,
	}
}

func (eh *eventHandler) StartEventHanlder() {
	eh.orderApproved()
	eh.orderRejected()
}
