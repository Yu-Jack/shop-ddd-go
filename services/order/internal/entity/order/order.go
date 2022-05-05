package order

import (
	"encoding/json"

	"github.com/Yu-Jack/shop-ddd-go/kit/dddcore"
)

type Order struct {
	ID           string          `json:"id" gorm:"primarykey"`
	ConsumerID   string          `json:"consumer_id"`
	Name         string          `json:"name"`
	State        string          `json:"state"`
	Amount       int             `json:"amount"`
	CreatedAt    int             `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt    int             `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DomainEvents []dddcore.Event `json:"-" gorm:"-:all"`
}

func NewOrder() *Order {
	o := &Order{}
	return o
}

func (o *Order) CreatedOrderEvent() {
	e := dddcore.NewEvent()
	e.RawData = []byte(o.ToJsonString())
	e.EventName = "OrderCreated"
	o.DomainEvents = append(o.DomainEvents, e)
}

func (o *Order) ToJsonString() string {
	data, _ := json.Marshal(o)
	return string(data)
}
