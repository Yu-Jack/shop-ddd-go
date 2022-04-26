package event

import "github.com/Yu-Jack/dddcore"

func NewOrderApproved(orderId string) []dddcore.Event {
	e := dddcore.NewEvent()
	e.EventName = "OrderApproved"
	e.RawData = []byte(orderId)
	return []dddcore.Event{e}
}

func NewOrderRejected(orderId string) []dddcore.Event {
	e := dddcore.NewEvent()
	e.EventName = "OrderRejected"
	e.RawData = []byte(orderId)
	return []dddcore.Event{e}
}
