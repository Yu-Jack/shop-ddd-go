package network

type CreateOrderItemReq struct {
	ConsumerID string `json:"user_id"`
	Name       string
	Amount     int
}
