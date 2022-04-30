package usecase

import orderEntity "github.com/Yu-Jack/shop-ddd-go-order/internal/entity/order"

type Repository interface {
	SaveOrderItem(oi orderEntity.OrderItem)
	GetAllOrderItems() []*orderEntity.OrderItem
}
