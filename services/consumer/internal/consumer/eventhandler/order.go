package eventhandler

import (
	"encoding/json"
)

type OrderCreatedEvent struct {
	OrderId    string `json:"id"`
	ConsumerID string `json:"consumer_id"`
	Amount     int    `json:"amount"`
}

func (eh *eventHandler) orderCreated() {
	go eh.eventBus.Subscribe("OrderCreated", func(value string) {
		e := OrderCreatedEvent{}
		json.Unmarshal([]byte(value), &e)
		eh.usecase.CheckOrder(e.OrderId, e.Amount, e.ConsumerID)
	})
}
