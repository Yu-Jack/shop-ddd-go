package user

type UserEvent interface {
	NewOrderApprovedEvent(eventName string, orderId string)
	NewOrderRejectedEvent(eventName string, orderId string)

	SubscribeOrderCreated(eventName string, cb func(value string))
}

type UserEventCacllback interface {
	OrderCreatedCallback(value string)
}
