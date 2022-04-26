package main

import (
	"github.com/Yu-Jack/dddcore"
	"github.com/Yu-Jack/shop-ddd-go-consumer/eventhandler"
	"github.com/Yu-Jack/shop-ddd-go-consumer/repository"
	"github.com/Yu-Jack/shop-ddd-go-consumer/request"
	"github.com/Yu-Jack/shop-ddd-go-consumer/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	eventBus := dddcore.NewEventBus()

	consumerRepo := repository.New()
	consumerUsecase := usecase.New(eventBus, consumerRepo)
	consumerEventHandler := eventhandler.New(consumerUsecase, eventBus)
	consumerEventHandler.StartEventHanlder()

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "consumer service is health",
		})
	})

	r.POST("/consumer", func(c *gin.Context) {
		var createConsumerReq request.CreateConsumerRequest
		if err := c.ShouldBindJSON(&createConsumerReq); err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		consumer := consumerUsecase.CreateConsumer(createConsumerReq)
		c.JSON(200, consumer)
	})

	r.GET("/consumers", func(c *gin.Context) {
		consumers := consumerUsecase.GetAllConsumers()
		c.JSON(200, gin.H{
			"consumers": consumers,
		})
	})

	r.Run(":8070")
}
