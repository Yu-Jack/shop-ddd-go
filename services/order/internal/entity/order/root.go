package order

var dbOrderItem []*OrderItem
var dbOrder []*Order

type repo struct{}

func New() *repo {
	return &repo{}
}

func (r *repo) SaveOrder(o *Order) {
	dbOrder = append(dbOrder, o)
}

func (r *repo) UpdateOrderState(orderid string, newState string) {
	for _, o := range dbOrder {
		if o.ID == orderid {
			o.State = newState
		}
	}
}

func (r *repo) FindOrderByIds(orderid string) *Order {
	for _, o := range dbOrder {
		if o.ID == orderid {
			return o
		}
	}
	return nil
}

func (r *repo) FindOrderItemByOrderId(orderId string) *OrderItem {
	for _, oi := range dbOrderItem {
		if oi.OrderID == orderId {
			return oi
		}
	}
	return nil
}

func (r *repo) FindOrderByConsumerId(consumerId string) *Order {
	for _, o := range dbOrder {
		if o.UserID == consumerId {
			return o
		}
	}
	return nil
}

func (r *repo) GetAllOrders() []*Order {
	return dbOrder
}

func (r *repo) GetAll() []*OrderItem {
	return dbOrderItem
}

func (r *repo) SaveOrderItem(oi OrderItem) {
	dbOrderItem = append(dbOrderItem, &oi)
}