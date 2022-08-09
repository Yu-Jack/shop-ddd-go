package order

import (
	domain "github.com/Yu-Jack/shop-ddd-go/internal/domain/order"
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

type Order interface {
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

type order struct {
	db *gorm.DB
}

func New(db *gorm.DB) Order {
	return &order{
		db: db,
	}
}

func (repoOrder repoOrder) transformDomain() domain.Order {
	return domain.Order{
		ID:         repoOrder.ID,
		ConsumerID: repoOrder.ConsumerID,
		Name:       repoOrder.Name,
		State:      repoOrder.State,
		Amount:     repoOrder.Amount,
		CreatedAt:  repoOrder.CreatedAt,
		UpdatedAt:  repoOrder.UpdatedAt,
	}
}

func (repoOrderItem repoOrderItem) transformDomain() domain.OrderItem {
	return domain.OrderItem{
		ID:        repoOrderItem.ID,
		OrderID:   repoOrderItem.OrderID,
		Name:      repoOrderItem.Name,
		Amount:    repoOrderItem.Amount,
		CreatedAt: repoOrderItem.CreatedAt,
		UpdatedAt: repoOrderItem.UpdatedAt,
	}
}
