package eventhandler

import (
	orderUc "github.com/Yu-Jack/shop-ddd-go-order/internal/order/usecase"
	"github.com/Yu-Jack/shop-ddd-go/kit/dddcore"
)

type eventHandler struct {
	orderUsecase orderUc.Usecase
	eventBus     *dddcore.EventBus
}

// No used for temporarily.
func New(orderUsecase orderUc.Usecase, eventBus *dddcore.EventBus) *eventHandler {
	return &eventHandler{
		orderUsecase: orderUsecase,
		eventBus:     eventBus,
	}
}
