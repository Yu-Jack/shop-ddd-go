package main

import (
	"github.com/Yu-Jack/dddcore"
	"github.com/Yu-Jack/shop-ddd-go-consumer/internal/consumer"
	"github.com/gin-gonic/gin"
)

func main() {
	eventBus := dddcore.NewEventBus()
	r := gin.Default()

	consumer.Register(r, eventBus)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "consumer service is health",
		})
	})

	r.Run(":8070")
}
