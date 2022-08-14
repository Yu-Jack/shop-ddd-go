package user

import (
	userUc "github.com/Yu-Jack/shop-ddd-go/internal/usecase/user"
	"github.com/Yu-Jack/shop-ddd-go/pkg/dddcore"
)

type user struct {
	userEventHandler userUc.UserEvent
	eventBus         *dddcore.EventBus
}

type User interface {
	StartEventHandler()
}

func New(userEventHandler userUc.UserEvent, eventBus *dddcore.EventBus) User {
	return &user{
		userEventHandler: userEventHandler,
		eventBus:         eventBus,
	}
}
