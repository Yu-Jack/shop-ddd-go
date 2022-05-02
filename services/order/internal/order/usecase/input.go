package usecase

type CreateOrderInput struct {
	ConsumerID string
	Name       string
}

type CheckoutOrderInput struct {
	ConsumerID string
}
