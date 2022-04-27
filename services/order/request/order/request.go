package order

type GetOrder struct {
	ID string `uri:"id"`
}

type CreateOrder struct {
	UserID string `json:"user_id"`
}
