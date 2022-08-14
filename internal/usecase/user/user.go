package user

import (
	userDomain "github.com/Yu-Jack/shop-ddd-go/internal/domain/user"
)

type CreateUserInput struct {
	FirstName string
	LastName  string
	Amount    int
}

type User interface {
	CheckOrder(orderId string, orderAmount int, userId string)
	CreateUser(input CreateUserInput) userDomain.User
	GetAllUsers() []userDomain.User
}

type user struct {
	repo  UserRepo
	event UserEvent
}

func New(repo UserRepo, event UserEvent) User {
	u := &user{
		repo:  repo,
		event: event,
	}

	u.event.SubscribeOrderCreated(userDomain.ORDER_CREATED_EVENT, u.OrderCreated)
	return u
}
