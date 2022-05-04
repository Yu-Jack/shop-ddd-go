package usecase

import (
	"fmt"

	"github.com/Yu-Jack/shop-ddd-go-consumer/internal/entity"
)

func (u *usecase) CheckOrder(orderId string, orderAmount int, consumerId string) {
	validToBuy := u.repo.DecreaseConsumerAmount(consumerId, orderAmount)

	if validToBuy {
		u.eventBus.Publish(entity.NewOrderApprovedEvent(orderId))
	} else {
		u.eventBus.Publish(entity.NewOrderRejectedEvent(orderId))
	}

	fmt.Println(orderId)
	fmt.Println("sent order final state")
}
