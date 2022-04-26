package usecase

import (
	"fmt"

	"github.com/Yu-Jack/shop-ddd-go-consumer/event"
)

func (u *usecase) CheckOrder(orderId string, orderAmount int, consumerId string) {
	validToBuy := u.repo.DecreaseConsumerAmount(consumerId, orderAmount)

	if validToBuy {
		u.eventBus.Publish(event.NewOrderApproved(orderId))
	} else {
		u.eventBus.Publish(event.NewOrderRejected(orderId))
	}

	fmt.Println(orderId)
	fmt.Println("sent order final state")
}
