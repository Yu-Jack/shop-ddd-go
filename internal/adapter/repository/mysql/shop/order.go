package shop

import (
	domain "github.com/Yu-Jack/shop-ddd-go/internal/domain/shop"
	shopUC "github.com/Yu-Jack/shop-ddd-go/internal/usecase/shop"
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

type order struct {
	db *gorm.DB
}

func New(db *gorm.DB) shopUC.ShopRepo {
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
