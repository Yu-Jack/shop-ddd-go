package usecase

import orderEntity "github.com/Yu-Jack/shop-ddd-go-order/internal/entity/order"

type Repository interface {
	SaveOrder(o *orderEntity.Order)
	UpdateOrderState(orderId string, newState string)
	FindOrderByIds(orderId string) *orderEntity.Order
	FindAvailableOrderByConsumerId(consumerId string) *orderEntity.Order
	GetAllOrders() []*orderEntity.Order
	FindOrderItemsByOrderId(orderId string) []*orderEntity.OrderItem
}
