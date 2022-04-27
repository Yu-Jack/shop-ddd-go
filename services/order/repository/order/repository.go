package order

import "github.com/Yu-Jack/shop-ddd-go-order/entity"

var db []*entity.Order

type repo struct{}

func New() *repo {
	return &repo{}
}
