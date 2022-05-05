package main

import (
	"github.com/Yu-Jack/shop-ddd-go-consumer/internal/consumer"
	"github.com/Yu-Jack/shop-ddd-go/kit/dddcore"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	eventBus := dddcore.NewEventBus()
	r := gin.Default()

	consumer.Register(r, eventBus, db)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "consumer service is health",
		})
	})

	r.Run(":8070")
}
