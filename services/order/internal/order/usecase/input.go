package usecase

type CreateOrderInput struct {
	UserID string
	Name   string
}

type CheckoutOrderInput struct {
	UserID string
}
