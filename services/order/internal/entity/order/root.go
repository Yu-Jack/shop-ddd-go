package order

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *repo {
	return &repo{
		db: db,
	}
}

func (r *repo) CreateOrder(o *Order) {
	r.db.Create(o)
}

func (r *repo) SaveOrder(o Order) {
	r.db.Save(o)
}

func (r *repo) UpdateOrderState(orderid string, newState string) error {
	result := r.db.Model(&Order{}).Where("id = ?", orderid).Where("state = ?", "CHECKOUT_PENDING").Update("state", newState)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repo) FindOrderById(orderid string) (o Order, err error) {
	result := r.db.Where("id = ?", orderid).Find(&o)

	if result.Error != nil {
		return o, result.Error
	}

	return o, nil
}

func (r *repo) FindTotalAmountByOrderId(orderId string) (amount int64, err error) {
	result := map[string]interface{}{}
	dbResult := r.db.Model(&OrderItem{}).Select("sum(amount) as total").Where("order_id = ?", orderId).Find(&result)

	if dbResult.Error != nil {
		return amount, dbResult.Error
	}

	return result["total"].(int64), nil
}

func (r *repo) FindAvailableOrderByConsumerId(consumerId string) (order Order, err error) {
	result := r.db.Where("state = ?", "PENDING").Where("consumer_id = ?", consumerId).First(&order)

	if result.Error != nil {
		return Order{}, result.Error
	}

	return order, nil
}

func (r *repo) GetAllOrders() (orders []Order, err error) {
	result := r.db.Find(&orders)
	if result.Error != nil {
		return orders, result.Error
	}
	return orders, nil
}

func (r *repo) GetAllOrderItemsByOrderId(orderId string) (ois []OrderItem, err error) {
	result := r.db.Where("order_id = ?", orderId).Find(&ois)
	if result.Error != nil {
		return ois, result.Error
	}
	return ois, nil
}

func (r *repo) CreateOrderItem(oi *OrderItem, consumerID string) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if oi.OrderID == "" {
			o := &Order{
				ID:         uuid.NewString(),
				ConsumerID: consumerID,
				Name:       fmt.Sprintf("OrderName - %s", uuid.NewString()),
				State:      "PENDING",
			}
			if err := tx.Create(o).Error; err != nil {
				return err
			}
			oi.OrderID = o.ID
		}
		if err := tx.Create(oi).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
