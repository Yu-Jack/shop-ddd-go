package user

import "encoding/json"

func (u *user) StartEventHandler() {
	u.orderCreated()
}

func (u *user) orderCreated() {
	type OrderCreatedEvent struct {
		OrderId    string `json:"id"`
		ConsumerID string `json:"consumer_id"`
		Amount     int    `json:"amount"`
	}

	go u.eventBus.Subscribe("OrderCreated", func(value string) {
		e := OrderCreatedEvent{}
		json.Unmarshal([]byte(value), &e)
		u.userEventHandler.CheckOrder(e.OrderId, e.Amount, e.ConsumerID)
	})
}
