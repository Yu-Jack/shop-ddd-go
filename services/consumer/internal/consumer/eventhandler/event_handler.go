package eventhandler

import (
	"github.com/Yu-Jack/dddcore"
	"github.com/Yu-Jack/shop-ddd-go-consumer/internal/consumer/usecase"
)

type eventHandler struct {
	usecase  usecase.Usecase
	eventBus *dddcore.EventBus
}

func New(usecase usecase.Usecase, eventBus *dddcore.EventBus) *eventHandler {
	return &eventHandler{
		usecase:  usecase,
		eventBus: eventBus,
	}
}

func (eh *eventHandler) StartEventHanlder() {
	eh.orderCreated()
}
