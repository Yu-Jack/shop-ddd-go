package network

type GetOrderReq struct {
	ID string `uri:"id"`
}

type CreateOrderReq struct {
	UserID string `json:"user_id"`
}
