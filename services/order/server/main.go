package main

import (
	"fmt"

	"github.com/Yu-Jack/dddcore"
	"github.com/Yu-Jack/shop-ddd-go-order/eventhandler"
	repoOrder "github.com/Yu-Jack/shop-ddd-go-order/repository/order"
	reqOrder "github.com/Yu-Jack/shop-ddd-go-order/request/order"
	ucOrder "github.com/Yu-Jack/shop-ddd-go-order/usecase/order"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	eventBus := dddcore.NewEventBus()

	orderRepo := repoOrder.New()
	orderUsecase := ucOrder.New(orderRepo, eventBus)
	orderEventHandler := eventhandler.New(orderUsecase, eventBus)
	orderEventHandler.StartEventHanlder()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "order service is health",
		})
	})

	r.GET("/order/:id", func(c *gin.Context) {
		var req reqOrder.GetOrder
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		o, _ := orderUsecase.FindOrderById(req.ID)
		c.JSON(200, o)
	})

	r.POST("/order", func(c *gin.Context) {
		var req reqOrder.CreateOrder
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		o, _ := orderUsecase.CreateOrder(ucOrder.CreateOrderInput{
			UserID: req.UserID,
			Name:   fmt.Sprintf("OrderName - %s", uuid.NewString()),
		})
		c.JSON(200, o)
	})

	r.GET("/orders", func(c *gin.Context) {
		orders, _ := orderUsecase.GetAllOrders()
		c.JSON(200, gin.H{
			"orders": orders,
		})
	})

	r.Run()
}
