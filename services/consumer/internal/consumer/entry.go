package consumer

import (
	"github.com/Yu-Jack/shop-ddd-go-consumer/internal/consumer/eventhandler"
	consumerNet "github.com/Yu-Jack/shop-ddd-go-consumer/internal/consumer/network"
	consumerUc "github.com/Yu-Jack/shop-ddd-go-consumer/internal/consumer/usecase"
	consumerEntity "github.com/Yu-Jack/shop-ddd-go-consumer/internal/entity"
	"github.com/Yu-Jack/shop-ddd-go/kit/dddcore"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(r *gin.Engine, eventBus *dddcore.EventBus, db *gorm.DB) {
	repo := consumerEntity.New(db)

	db.AutoMigrate(&consumerEntity.Consumer{})

	consumerUsecase := consumerUc.New(eventBus, repo)
	eventhandler.New(consumerUsecase, eventBus).StartEventHanlder()
	consumerNet.New(r, consumerUsecase).Route()
}
