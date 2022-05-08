package order

import (
	orderEntity "github.com/Yu-Jack/shop-ddd-go-order/internal/entity/order"
	orderCtrl "github.com/Yu-Jack/shop-ddd-go-order/internal/order/controller"
	orderEventHandler "github.com/Yu-Jack/shop-ddd-go-order/internal/order/eventhandler"
	orderSaga "github.com/Yu-Jack/shop-ddd-go-order/internal/order/saga"
	orderUc "github.com/Yu-Jack/shop-ddd-go-order/internal/order/usecase"
	"github.com/Yu-Jack/shop-ddd-go/kit/dddcore"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(r *gin.Engine, eventBus *dddcore.EventBus, db *gorm.DB) {
	repo := orderEntity.New(db)

	orderUsecase := orderUc.New(repo, eventBus)
	orderSa := orderSaga.NewSaga(eventBus, orderUsecase)

	orderEventHandler := orderEventHandler.New(orderUsecase, eventBus)
	orderEventHandler.StartEventHanlder()
	orderCtrl.New(r, orderUsecase, orderSa).Route()
}
