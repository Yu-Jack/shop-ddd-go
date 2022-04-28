package repository

import "github.com/Yu-Jack/shop-ddd-go-order/internal/order/entity"

var db []*entity.Order

type repo struct{}

func New() *repo {
	return &repo{}
}
