package request

type CreateConsumerRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Amount    int    `json:"amount"`
}
