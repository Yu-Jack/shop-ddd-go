package usecase

import "github.com/Yu-Jack/shop-ddd-go-order/internal/order/entity"

type Repository interface {
	Save(o *entity.Order)
	UpdateOrderState(orderId string, newState string)
	FindOrderByIds(orderId string) *entity.Order
	GetAllOrders() []*entity.Order
}
