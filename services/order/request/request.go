package request

type GetOrderRequest struct {
	ID string `uri:"id"`
}

type CreateOrderRequest struct {
	UserID string `json:"user_id"`
}
