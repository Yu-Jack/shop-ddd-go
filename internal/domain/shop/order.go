package shop

import (
	"encoding/json"

	"github.com/Yu-Jack/shop-ddd-go/pkg/dddcore"
)

type Order struct {
	ID         string `json:"id"`
	ConsumerID string `json:"consumer_id"`
	Name       string `json:"name"`
	State      string `json:"state"`
	Amount     int    `json:"amount"`
	CreatedAt  int    `json:"created_at"`
	UpdatedAt  int    `json:"updated_at"`
}

func (o Order) CreatedOrderEvent() []dddcore.Event {
	e := dddcore.NewEvent()
	e.RawData = []byte(o.ToJsonString())
	e.EventName = "OrderCreated"
	return []dddcore.Event{e}
}

func (o *Order) ToJsonString() string {
	data, _ := json.Marshal(o)
	return string(data)
}
