package user

type User struct {
	ID        string
	FirstName string
	LastName  string
	Amount    int
	CreatedAt int
	UpdatedAt int
}

var ORDER_CREATED_EVENT = "OrderCreated"
var ORDER_APPROVED_EVENT = "OrderApproved"
var ORDER_REJECTED_EVENT = "OrderRejected"
