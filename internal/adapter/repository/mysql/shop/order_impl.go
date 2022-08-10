package shop

import (
	"fmt"

	"github.com/Yu-Jack/shop-ddd-go/internal/domain/shop"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *repo) CreateOrder(o *shop.Order) {
	r.db.Create(o)
}

func (r *repo) SaveOrder(o shop.Order) {
	r.db.Save(o)
}

func (r *repo) UpdateOrderState(orderid string, newState string) error {
	result := r.db.Model(&repoOrder{}).Where("id = ?", orderid).Where("state = ?", "CHECKOUT_PENDING").Update("state", newState)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repo) FindOrderById(orderid string) (shop.Order, error) {
	repoO := repoOrder{}
	result := r.db.Where("id = ?", orderid).Find(&repoO)

	if result.Error != nil {
		return shop.Order{}, result.Error
	}

	return repoO.transformDomain(), nil
}

func (r *repo) FindTotalAmountByOrderId(orderId string) (amount int64, err error) {
	result := map[string]interface{}{}
	dbResult := r.db.Model(&repoOrderItem{}).Select("sum(amount) as total").Where("order_id = ?", orderId).Find(&result)

	if dbResult.Error != nil {
		return amount, dbResult.Error
	}

	return result["total"].(int64), nil
}

func (r *repo) FindAvailableOrderByConsumerId(consumerId string) (order shop.Order, err error) {
	repoO := repoOrder{}
	result := r.db.Where("state = ?", "PENDING").Where("consumer_id = ?", consumerId).First(&repoO)

	if result.Error != nil {
		return shop.Order{}, result.Error
	}

	return repoO.transformDomain(), nil
}

func (r *repo) GetAllOrders() (orders []shop.Order, err error) {
	repoOs := []repoOrder{}
	result := r.db.Find(&repoOs)
	if result.Error != nil {
		return []shop.Order{}, result.Error
	}

	for _, repoO := range repoOs {
		orders = append(orders, repoO.transformDomain())
	}

	return orders, nil
}

func (r *repo) GetAllOrderItemsByOrderId(orderId string) (ois []shop.OrderItem, err error) {
	repoOrderItems := []repoOrderItem{}
	result := r.db.Where("order_id = ?", orderId).Find(&repoOrderItems)
	if result.Error != nil {
		return []shop.OrderItem{}, result.Error
	}

	for _, repoOrderItem := range repoOrderItems {
		ois = append(ois, repoOrderItem.transformDomain())
	}

	return ois, nil
}

func (r *repo) CreateOrderItem(oi *shop.OrderItem, consumerID string) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if oi.OrderID == "" {
			o := &repoOrder{
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
