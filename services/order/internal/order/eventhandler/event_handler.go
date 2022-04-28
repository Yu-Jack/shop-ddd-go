package eventhandler

import (
	"github.com/Yu-Jack/dddcore"
	orderUc "github.com/Yu-Jack/shop-ddd-go-order/internal/order/usecase"
)

type eventHandler struct {
	orderUsecase orderUc.Usecase
	eventBus     *dddcore.EventBus
}

func New(orderUsecase orderUc.Usecase, eventBus *dddcore.EventBus) *eventHandler {
	return &eventHandler{
		orderUsecase: orderUsecase,
		eventBus:     eventBus,
	}
}

func (eh *eventHandler) StartEventHanlder() {
	eh.orderApproved()
	eh.orderRejected()
}
