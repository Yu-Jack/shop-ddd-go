package network

type CreateOrderItemReq struct {
	ConsumerID string `json:"user_id"`
	Name       string `json:"name"`
	Amount     int    `json:"amount"`
}

type GetOrderItemReq struct {
	OrderID string `uri:"order_id"`
}
