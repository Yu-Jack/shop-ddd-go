package user

import (
	"fmt"

	"github.com/Yu-Jack/shop-ddd-go/pkg/dddcore"
)

type UserEvent interface {
	CheckOrder(orderId string, orderAmount int, userId string)
}

func (u *user) newOrderApprovedEvent(orderId string) []dddcore.Event {
	e := dddcore.NewEvent()
	e.EventName = fmt.Sprintf("%s-%s", "OrderApproved", orderId)
	e.RawData = []byte(orderId)
	return []dddcore.Event{e}
}

func (u *user) newOrderRejectedEvent(orderId string) []dddcore.Event {
	e := dddcore.NewEvent()
	e.EventName = fmt.Sprintf("%s-%s", "OrderRejected", orderId)
	e.RawData = []byte(orderId)
	return []dddcore.Event{e}
}
