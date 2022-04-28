package order

import (
	"github.com/Yu-Jack/dddcore"
	orderEventHandler "github.com/Yu-Jack/shop-ddd-go-order/internal/order/eventhandler"
	orderNet "github.com/Yu-Jack/shop-ddd-go-order/internal/order/network"
	orderRepo "github.com/Yu-Jack/shop-ddd-go-order/internal/order/repository"
	orderUc "github.com/Yu-Jack/shop-ddd-go-order/internal/order/usecase"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, eventBus *dddcore.EventBus) {
	oRepo := orderRepo.New()
	orderUsecase := orderUc.New(oRepo, eventBus)
	orderEventHandler := orderEventHandler.New(orderUsecase, eventBus)
	orderEventHandler.StartEventHanlder()
	oNet := orderNet.New(r, orderUsecase)
	oNet.Route()
}
