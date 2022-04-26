package main

import (
	"fmt"

	"github.com/Yu-Jack/dddcore"
	"github.com/Yu-Jack/shop-ddd-go-order/eventhandler"
	"github.com/Yu-Jack/shop-ddd-go-order/repository"
	"github.com/Yu-Jack/shop-ddd-go-order/request"
	"github.com/Yu-Jack/shop-ddd-go-order/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	eventBus := dddcore.NewEventBus()

	orderRepo := repository.New()
	orderUsecase := usecase.New(orderRepo, eventBus)
	orderEventHandler := eventhandler.New(orderUsecase, eventBus)
	orderEventHandler.StartEventHanlder()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "order service is health",
		})
	})

	r.GET("/order/:id", func(c *gin.Context) {
		var getOrderReq request.GetOrderRequest
		if err := c.ShouldBindUri(&getOrderReq); err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		o, _ := orderUsecase.FindOrderById(getOrderReq.ID)
		c.JSON(200, o)
	})

	r.POST("/order", func(c *gin.Context) {
		var createOrderReq request.CreateOrderRequest
		if err := c.ShouldBindJSON(&createOrderReq); err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		o, _ := orderUsecase.CreateOrder(usecase.CreateOrderInput{
			UserID: createOrderReq.UserID,
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
