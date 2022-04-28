package order

import (
	"github.com/Yu-Jack/dddcore"

	orderEntity "github.com/Yu-Jack/shop-ddd-go-order/internal/entity/order"
	orderUc "github.com/Yu-Jack/shop-ddd-go-order/internal/order/usecase"
	orderItemNet "github.com/Yu-Jack/shop-ddd-go-order/internal/order_item/network"
	orderItemUc "github.com/Yu-Jack/shop-ddd-go-order/internal/order_item/usecase"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, eventBus *dddcore.EventBus) {
	repo := orderEntity.New()

	orderUsecase := orderUc.New(repo, eventBus)
	orderItemUsecase := orderItemUc.New(repo, eventBus)

	orderItemNet.New(r, orderUsecase, orderItemUsecase).Route()
}
