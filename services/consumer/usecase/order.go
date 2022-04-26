package usecase

import "fmt"

func (u *usecase) CheckOrder(orderId string, orderAmount int, consumerId string) {
	validToBuy := u.repo.DecreaseConsumerAmount(consumerId, orderAmount)

	if validToBuy {
		u.eventBus.Publish("OrderApproved", orderId)
	} else {
		u.eventBus.Publish("OrderRejected", orderId)
	}

	fmt.Println(orderId)
	fmt.Println("sent order final state")
}
