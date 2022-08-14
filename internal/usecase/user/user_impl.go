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
	u.repo.DecreaseUserAmount(userId, orderAmount)

	// TODO: should send event
	// if err == nil {
	// 	u.eventBus.Publish(entity.NewOrderApprovedEvent(orderId))
	// } else {
	// 	u.eventBus.Publish(entity.NewOrderRejectedEvent(orderId))
	// }

	fmt.Println(orderId)
	fmt.Println("sent order final state")
}
