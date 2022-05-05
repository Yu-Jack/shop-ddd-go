package entity

import "github.com/Yu-Jack/shop-ddd-go/kit/dddcore"

func NewOrderApprovedEvent(orderId string) []dddcore.Event {
	e := dddcore.NewEvent()
	e.EventName = "OrderApproved"
	e.RawData = []byte(orderId)
	return []dddcore.Event{e}
}

func NewOrderRejectedEvent(orderId string) []dddcore.Event {
	e := dddcore.NewEvent()
	e.EventName = "OrderRejected"
	e.RawData = []byte(orderId)
	return []dddcore.Event{e}
}
