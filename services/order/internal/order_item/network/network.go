package network

import (
	"fmt"

	orderUc "github.com/Yu-Jack/shop-ddd-go-order/internal/order/usecase"
	orderItemUc "github.com/Yu-Jack/shop-ddd-go-order/internal/order_item/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type net struct {
	r           *gin.Engine
	orderUc     orderUc.Usecase
	orderItemUc orderItemUc.Usecase
}

type Net interface {
	Route()
}

func New(r *gin.Engine, orderUc orderUc.Usecase, orderItemUc orderItemUc.Usecase) Net {
	return &net{
		r:           r,
		orderUc:     orderUc,
		orderItemUc: orderItemUc,
	}
}

func (n *net) Route() {
	n.r.POST("/order/item", n.createOrderItem)
	n.r.GET("/order/item/:order_id", n.getOrderItems)
}

func (n *net) createOrderItem(c *gin.Context) {
	var req CreateOrderItemReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	o, _ := n.orderUc.FindOrderByConsumerId(req.ConsumerID)
	if o == nil {
		o, _ = n.orderUc.CreateOrder(orderUc.CreateOrderInput{
			UserID: req.ConsumerID,
			Name:   fmt.Sprintf("OrderName - %s", uuid.NewString()),
		})
	}

	oi, _ := n.orderItemUc.CreateOrderItem(orderItemUc.CreateOrderItemInput{
		OrderID: o.ID,
		Name:    req.Name,
		Amount:  req.Amount,
	})
	c.JSON(200, oi)
}

func (n *net) getOrderItems(c *gin.Context) {
	var req GetOrderItemReq
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	ois, _ := n.orderItemUc.GetOrderItems(req.OrderID)
	c.JSON(200, ois)
}