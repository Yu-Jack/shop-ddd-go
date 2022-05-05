package order

import (
	"github.com/Yu-Jack/dddcore"
	orderEntity "github.com/Yu-Jack/shop-ddd-go-order/internal/entity/order"
	orderEventHandler "github.com/Yu-Jack/shop-ddd-go-order/internal/order/eventhandler"
	orderNet "github.com/Yu-Jack/shop-ddd-go-order/internal/order/network"
	orderUc "github.com/Yu-Jack/shop-ddd-go-order/internal/order/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(r *gin.Engine, eventBus *dddcore.EventBus, db *gorm.DB) {
	repo := orderEntity.New(db)
	orderUsecase := orderUc.New(repo, eventBus)
	orderEventHandler := orderEventHandler.New(orderUsecase, eventBus)
	orderEventHandler.StartEventHanlder()
	oNet := orderNet.New(r, orderUsecase)
	oNet.Route()
}
