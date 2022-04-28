package order

import (
	"github.com/Yu-Jack/dddcore"
	orderEntity "github.com/Yu-Jack/shop-ddd-go-order/internal/entity/order"
	orderEventHandler "github.com/Yu-Jack/shop-ddd-go-order/internal/order/eventhandler"
	orderNet "github.com/Yu-Jack/shop-ddd-go-order/internal/order/network"
	orderUc "github.com/Yu-Jack/shop-ddd-go-order/internal/order/usecase"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, eventBus *dddcore.EventBus) {
	repo := orderEntity.New()
	orderUsecase := orderUc.New(repo, eventBus)
	orderEventHandler := orderEventHandler.New(orderUsecase, eventBus)
	orderEventHandler.StartEventHanlder()
	oNet := orderNet.New(r, orderUsecase)
	oNet.Route()
}