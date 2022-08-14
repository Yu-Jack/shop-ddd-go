package user

import (
	"fmt"

	userUc "github.com/Yu-Jack/shop-ddd-go/internal/usecase/user"
	"github.com/Yu-Jack/shop-ddd-go/pkg/dddcore"
)

type event struct {
	eventBus *dddcore.EventBus
}

type Event interface {
	userUc.UserEvent
}

func New(eventBus *dddcore.EventBus) userUc.UserEvent {
	return &event{
		eventBus: eventBus,
	}
}

func (event *event) NewOrderApprovedEvent(eventName string, orderId string) {
	domainEvent := dddcore.NewEvent()
	domainEvent.EventName = fmt.Sprintf("%s-%s", eventName, orderId)
	domainEvent.RawData = []byte(orderId)

	event.eventBus.Publish([]dddcore.Event{domainEvent})
}

func (event *event) NewOrderRejectedEvent(eventName string, orderId string) {
	domainEvent := dddcore.NewEvent()
	domainEvent.EventName = fmt.Sprintf("%s-%s", eventName, orderId)
	domainEvent.RawData = []byte(orderId)

	event.eventBus.Publish([]dddcore.Event{domainEvent})
}

func (event *event) SubscribeOrderCreated(eventName string, callback func(value string)) {
	go event.eventBus.Subscribe(eventName, callback)
}
