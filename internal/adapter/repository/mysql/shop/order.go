package shop

import (
	"github.com/Yu-Jack/shop-ddd-go/internal/domain/shop"
	"gorm.io/gorm"
)

type repoOrder struct {
	ID         string `json:"id" gorm:"primarykey"`
	ConsumerID string `json:"consumer_id"`
	Name       string `json:"name"`
	State      string `json:"state"`
	Amount     int    `json:"amount"`
	CreatedAt  int    `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt  int    `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

type repoOrderItem struct {
	ID        string `json:"id" gorm:"primarykey"`
	OrderID   string `json:"order_id"`
	Name      string `json:"name"`
	Amount    int    `json:"amount"`
	CreatedAt int    `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int    `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

type Repo interface {
	CreateOrder(o *shop.Order)
	SaveOrder(o shop.Order)
	UpdateOrderState(orderId string, newState string) error
	FindOrderById(orderId string) (shop.Order, error)
	FindAvailableOrderByConsumerId(consumerId string) (shop.Order, error)
	GetAllOrders() ([]shop.Order, error)
	FindTotalAmountByOrderId(orderId string) (amount int64, err error)
	CreateOrderItem(oi *shop.OrderItem, consumerID string) error
	GetAllOrderItemsByOrderId(orderId string) ([]shop.OrderItem, error)
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{
		db: db,
	}
}

func (repoOrder repoOrder) transformDomain() shop.Order {
	return shop.Order{
		ID:         repoOrder.ID,
		ConsumerID: repoOrder.ConsumerID,
		Name:       repoOrder.Name,
		State:      repoOrder.State,
		Amount:     repoOrder.Amount,
		CreatedAt:  repoOrder.CreatedAt,
		UpdatedAt:  repoOrder.UpdatedAt,
	}
}

func (repoOrderItem repoOrderItem) transformDomain() shop.OrderItem {
	return shop.OrderItem{
		ID:        repoOrderItem.ID,
		OrderID:   repoOrderItem.OrderID,
		Name:      repoOrderItem.Name,
		Amount:    repoOrderItem.Amount,
		CreatedAt: repoOrderItem.CreatedAt,
		UpdatedAt: repoOrderItem.UpdatedAt,
	}
}
