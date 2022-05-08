package entity

import (
	"fmt"

	"github.com/Yu-Jack/shop-ddd-go/kit/dddcore"
)

func NewOrderApprovedEvent(orderId string) []dddcore.Event {
	e := dddcore.NewEvent()
	e.EventName = fmt.Sprintf("%s-%s", "OrderApproved", orderId)
	e.RawData = []byte(orderId)
	return []dddcore.Event{e}
}

func NewOrderRejectedEvent(orderId string) []dddcore.Event {
	e := dddcore.NewEvent()
	e.EventName = fmt.Sprintf("%s-%s", "OrderRejected", orderId)
	e.RawData = []byte(orderId)
	return []dddcore.Event{e}
}
