package order

import (
	"github.com/Yu-Jack/shop-ddd-go/kit/dddcore"
	"gorm.io/gorm"

	orderEntity "github.com/Yu-Jack/shop-ddd-go-order/internal/entity/order"
	orderUc "github.com/Yu-Jack/shop-ddd-go-order/internal/order/usecase"
	orderItemNet "github.com/Yu-Jack/shop-ddd-go-order/internal/order_item/network"
	orderItemUc "github.com/Yu-Jack/shop-ddd-go-order/internal/order_item/usecase"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, eventBus *dddcore.EventBus, db *gorm.DB) {
	repo := orderEntity.New(db)

	db.AutoMigrate(&orderEntity.Order{})
	db.AutoMigrate(&orderEntity.OrderItem{})

	orderUsecase := orderUc.New(repo, eventBus)
	orderItemUsecase := orderItemUc.New(repo, eventBus)

	orderItemNet.New(r, orderUsecase, orderItemUsecase).Route()
}
