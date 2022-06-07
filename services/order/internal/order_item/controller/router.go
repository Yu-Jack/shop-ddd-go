package controller

import (
	orderUc "github.com/Yu-Jack/shop-ddd-go-order/internal/order/usecase"
	orderItemUc "github.com/Yu-Jack/shop-ddd-go-order/internal/order_item/usecase"
	"github.com/gin-gonic/gin"
)

type ctrl struct {
	r           *gin.Engine
	orderUc     orderUc.Usecase
	orderItemUc orderItemUc.Usecase
}

type Ctrl interface {
	Route()
}

func New(r *gin.Engine, orderUc orderUc.Usecase, orderItemUc orderItemUc.Usecase) Ctrl {
	return &ctrl{
		r:           r,
		orderUc:     orderUc,
		orderItemUc: orderItemUc,
	}
}

func (n *ctrl) Route() {
	n.r.POST("/order/item", n.createOrderItem)
	n.r.GET("/order/item/:order_id", n.getOrderItems)
}

func (n *ctrl) createOrderItem(c *gin.Context) {
	var req CreateOrderItemReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	o, _ := n.orderUc.FindAvailableOrderByConsumerId(c, req.ConsumerID)
	oi, _ := n.orderItemUc.CreateOrderItem(c, orderItemUc.CreateOrderItemInput{
		OrderID:    o.ID,
		Name:       req.Name,
		Amount:     req.Amount,
		ConsumerID: req.ConsumerID,
	})
	c.JSON(200, oi)
}

func (n *ctrl) getOrderItems(c *gin.Context) {
	var req GetOrderItemReq
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	ois, _ := n.orderItemUc.GetOrderItems(c, req.OrderID)
	c.JSON(200, ois)
}
