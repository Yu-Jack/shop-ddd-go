package main

import (
	"github.com/Yu-Jack/dddcore"
	orderEntry "github.com/Yu-Jack/shop-ddd-go-order/internal/order"
	orderItemEntry "github.com/Yu-Jack/shop-ddd-go-order/internal/order_item"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	eventBus := dddcore.NewEventBus()

	orderEntry.Register(r, eventBus)
	orderItemEntry.Register(r, eventBus)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "order service is health",
		})
	})

	r.Run()
}
