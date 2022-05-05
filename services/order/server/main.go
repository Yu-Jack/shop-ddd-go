package main

import (
	"github.com/Yu-Jack/dddcore"
	orderEntry "github.com/Yu-Jack/shop-ddd-go-order/internal/order"
	orderItemEntry "github.com/Yu-Jack/shop-ddd-go-order/internal/order_item"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	r := gin.Default()
	eventBus := dddcore.NewEventBus()

	orderEntry.Register(r, eventBus, db)
	orderItemEntry.Register(r, eventBus, db)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "order service is health",
		})
	})

	r.Run()
}
