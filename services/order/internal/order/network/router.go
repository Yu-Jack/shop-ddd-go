package network

import (
	"github.com/Yu-Jack/shop-ddd-go-order/internal/order/error_code"
	orderUc "github.com/Yu-Jack/shop-ddd-go-order/internal/order/usecase"
	"github.com/gin-gonic/gin"
)

type net struct {
	r       *gin.Engine
	orderUc orderUc.Usecase
}

type Net interface {
	Route()
}

func New(r *gin.Engine, orderUc orderUc.Usecase) Net {
	return &net{
		r:       r,
		orderUc: orderUc,
	}
}

func (n *net) Route() {
	n.r.POST("/order/checkout", n.checkoutOrder)
	n.r.GET("/order/:id", n.getOrder)
	n.r.GET("/orders", n.getOrders)
}

func (n *net) checkoutOrder(c *gin.Context) {
	var req CreateOrderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, error_code.New(error_code.REQUEST_BODY_IS_INVALID))
		return
	}
	o, err := n.orderUc.CheckoutOrder(orderUc.CheckoutOrderInput{
		ConsumerID: req.ConsumerID,
	})

	if err != nil {
		c.JSON(200, error_code.New(error_code.ORDER_IS_NOT_AVAILABLE))
		return
	}

	c.JSON(200, o)
}

func (n *net) getOrder(c *gin.Context) {
	var req GetOrderReq
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, error_code.New(error_code.REQUEST_URI_IS_INVALID))
		return
	}

	o, _ := n.orderUc.FindOrderById(req.ID)

	c.JSON(200, o)
}

func (n *net) getOrders(c *gin.Context) {
	orders, _ := n.orderUc.GetAllOrders()
	c.JSON(200, gin.H{
		"orders": orders,
	})
}
