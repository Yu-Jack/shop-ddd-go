package usecase

import orderEntity "github.com/Yu-Jack/shop-ddd-go-order/internal/entity/order"

type Repository interface {
	CreateOrderItem(oi *orderEntity.OrderItem)
	GetAllOrderItemsByOrderId(orderId string) ([]orderEntity.OrderItem, error)
}
