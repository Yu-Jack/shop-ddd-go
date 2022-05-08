package controller

type CreateOrderItemReq struct {
	ConsumerID string `json:"consumer_id"`
	Name       string `json:"name"`
	Amount     int    `json:"amount"`
}

type GetOrderItemReq struct {
	OrderID string `uri:"order_id"`
}
