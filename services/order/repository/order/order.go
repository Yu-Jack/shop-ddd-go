package order

import "github.com/Yu-Jack/shop-ddd-go-order/entity"

func (r *repo) Save(o *entity.Order) {
	db = append(db, o)
}

func (r *repo) UpdateOrderState(orderid string, newState string) {
	for _, o := range db {
		if o.ID == orderid {
			o.State = newState
		}
	}
}

func (r *repo) FindOrderByIds(orderid string) *entity.Order {
	for _, o := range db {
		if o.ID == orderid {
			return o
		}
	}
	return nil
}

func (r *repo) GetAllOrders() []*entity.Order {
	return db
}
