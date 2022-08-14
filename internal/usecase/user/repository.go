package user

import (
	userDomain "github.com/Yu-Jack/shop-ddd-go/internal/domain/user"
)

type UserRepo interface {
	DecreaseUserAmount(userId string, orderAmount int) error
	CreateUser(c *userDomain.User) error
	GetAllUsers() ([]userDomain.User, error)
}
