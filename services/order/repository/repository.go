package repository

import "github.com/Yu-Jack/shop-ddd-go-order/entity"

var db []*entity.Order

type repo struct{}

type Repository interface {
	Save(o entity.Order)
	UpdateOrderState(orderId string, newState string)
	FindOrderByIds(orderId string) *entity.Order
	GetAllOrders() []*entity.Order
}

func New() Repository {
	return &repo{}
}
