package usecase

type CreateOrderItemInput struct {
	Name    string
	Amount  int
	OrderID string
}
