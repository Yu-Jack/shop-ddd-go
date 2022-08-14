package user

import (
	userDomain "github.com/Yu-Jack/shop-ddd-go/internal/domain/user"
	"github.com/Yu-Jack/shop-ddd-go/pkg/dddcore"
)

type CreateUserInput struct {
	FirstName string
	LastName  string
	Amount    int
}

type User interface {
	CreateUser(input CreateUserInput) userDomain.User
	GetAllUsers() []userDomain.User
}

// TODO: renaming
type UserAll interface {
	User
	UserEvent
}

type user struct {
	repo     UserRepo
	eventBus *dddcore.EventBus
}

func New(repo UserRepo, eventBus *dddcore.EventBus) UserAll {
	return &user{
		repo:     repo,
		eventBus: eventBus,
	}
}
