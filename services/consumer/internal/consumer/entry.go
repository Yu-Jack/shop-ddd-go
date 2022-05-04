package consumer

import (
	"github.com/Yu-Jack/dddcore"
	"github.com/Yu-Jack/shop-ddd-go-consumer/internal/consumer/eventhandler"
	consumerNet "github.com/Yu-Jack/shop-ddd-go-consumer/internal/consumer/network"
	consumerUc "github.com/Yu-Jack/shop-ddd-go-consumer/internal/consumer/usecase"
	"github.com/Yu-Jack/shop-ddd-go-consumer/internal/entity"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, eventBus *dddcore.EventBus) {
	repo := entity.New()
	consumerUsecase := consumerUc.New(eventBus, repo)
	eventhandler.New(consumerUsecase, eventBus).StartEventHanlder()
	consumerNet.New(r, consumerUsecase).Route()
}
