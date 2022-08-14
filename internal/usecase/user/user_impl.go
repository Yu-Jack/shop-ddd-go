package user

import (
	"fmt"

	userDomain "github.com/Yu-Jack/shop-ddd-go/internal/domain/user"
	"github.com/google/uuid"
)

func (u *user) CreateUser(input CreateUserInput) userDomain.User {
	c := &userDomain.User{
		ID:        uuid.NewString(),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Amount:    input.Amount,
	}
	u.repo.CreateUser(c)
	return *c
}

func (u *user) GetAllUsers() []userDomain.User {
	cs, _ := u.repo.GetAllUsers()
	return cs
}

// TODO: should receive event to do check order
func (u *user) CheckOrder(orderId string, orderAmount int, userId string) {
	err := u.repo.DecreaseUserAmount(userId, orderAmount)

	if err == nil {
		u.event.NewOrderApprovedEvent(userDomain.ORDER_APPROVED_EVENT, orderId)
	} else {
		u.event.NewOrderRejectedEvent(userDomain.ORDER_REJECTED_EVENT, orderId)
	}

	fmt.Println(orderId)
	fmt.Println("sent order final state")
}
