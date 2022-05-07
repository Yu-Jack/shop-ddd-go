package network

type GetOrderReq struct {
	ID string `uri:"id"`
}

type CreateOrderReq struct {
	ConsumerID string `json:"consumer_id"`
}
