package controller

import (
	"github.com/Yu-Jack/shop-ddd-go-order/internal/order/error_code"
	orderUc "github.com/Yu-Jack/shop-ddd-go-order/internal/order/usecase"
	logger "github.com/Yu-Jack/shop-ddd-go/kit/logger"
	"github.com/gin-gonic/gin"
)

type ctrl struct {
	r       *gin.Engine
	orderUc orderUc.Usecase
}

type Ctrl interface {
	Route()
}

func New(r *gin.Engine, orderUc orderUc.Usecase) Ctrl {
	return &ctrl{
		r:       r,
		orderUc: orderUc,
	}
}

func (n *ctrl) Route() {
	n.r.POST("/order/checkout", n.checkoutOrder)
	n.r.GET("/order/:id", n.getOrder)
	n.r.GET("/orders", n.getOrders)
}

func (n *ctrl) checkoutOrder(c *gin.Context) {
	var req CreateOrderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, error_code.New(error_code.REQUEST_BODY_IS_INVALID))
		return
	}
	o, err := n.orderUc.CheckoutOrder(c, orderUc.CheckoutOrderInput{
		ConsumerID: req.ConsumerID,
	})

	if err != nil {
		logger.Log(c, "err", err)
		c.JSON(200, error_code.New(error_code.ORDER_IS_NOT_AVAILABLE))
		return
	}

	c.JSON(200, o)
}

func (n *ctrl) getOrder(c *gin.Context) {
	var req GetOrderReq
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, error_code.New(error_code.REQUEST_URI_IS_INVALID))
		return
	}

	o, _ := n.orderUc.FindOrderById(c, req.ID)

	c.JSON(200, o)
}

func (n *ctrl) getOrders(c *gin.Context) {
	orders, _ := n.orderUc.GetAllOrders(c)
	c.JSON(200, gin.H{
		"orders": orders,
	})
}
