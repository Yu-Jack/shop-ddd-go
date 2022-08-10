package shop

import domain "github.com/Yu-Jack/shop-ddd-go/internal/domain/shop"

type ShopRepo interface {
	CreateOrder(o *domain.Order)
	SaveOrder(o domain.Order)
	UpdateOrderState(orderId string, newState string) error
	FindOrderById(orderId string) (domain.Order, error)
	FindAvailableOrderByConsumerId(consumerId string) (domain.Order, error)
	GetAllOrders() ([]domain.Order, error)
	FindTotalAmountByOrderId(orderId string) (amount int64, err error)
	CreateOrderItem(oi *domain.OrderItem, consumerID string) error
	GetAllOrderItemsByOrderId(orderId string) ([]domain.OrderItem, error)
}
