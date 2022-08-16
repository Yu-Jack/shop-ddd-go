package user

import "encoding/json"

func (user *user) OrderCreatedCallback(value string) {
	type OrderCreatedEvent struct {
		OrderId    string `json:"id"`
		ConsumerID string `json:"consumer_id"`
		Amount     int    `json:"amount"`
	}

	e := OrderCreatedEvent{}
	json.Unmarshal([]byte(value), &e)
	user.CheckOrder(e.OrderId, e.Amount, e.ConsumerID)
}
