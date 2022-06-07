package usecase

import orderEntity "github.com/Yu-Jack/shop-ddd-go-order/internal/entity/order"

type Repository interface {
	CreateOrder(o *orderEntity.Order)
	SaveOrder(o orderEntity.Order)
	UpdateOrderState(orderId string, newState string) error
	FindOrderById(orderId string) (orderEntity.Order, error)
	FindAvailableOrderByConsumerId(consumerId string) (orderEntity.Order, error)
	GetAllOrders() ([]orderEntity.Order, error)
	FindTotalAmountByOrderId(orderId string) (amount int64, err error)
	CreateOrderItem(oi *orderEntity.OrderItem, consumerID string) error
	GetAllOrderItemsByOrderId(orderId string) ([]orderEntity.OrderItem, error)
}
