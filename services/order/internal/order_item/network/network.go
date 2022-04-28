package network

import (
	orderUc "github.com/Yu-Jack/shop-ddd-go-order/internal/order/usecase"
	orderItemUc "github.com/Yu-Jack/shop-ddd-go-order/internal/order_item/usecase"
	"github.com/gin-gonic/gin"
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
	// TODO: When finish order logic, then open it.
	// n.r.POST("/order/item", n.createOrderItem)
}

func (n *net) createOrderItem(c *gin.Context) {
	var req CreateOrderItemReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	o, _ := n.orderUc.FindOrderByConsumerId(req.ConsumerID)
	oi, _ := n.orderItemUc.CreateOrderItem(orderItemUc.CreateOrderItemInput{
		OrderID: o.ID,
		Name:    req.Name,
		Amount:  req.Amount,
	})
	c.JSON(200, oi)
}
