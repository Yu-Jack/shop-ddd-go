package controller

import (
	"github.com/Yu-Jack/shop-ddd-go-order/internal/order/error_code"
	orderUc "github.com/Yu-Jack/shop-ddd-go-order/internal/order/usecase"
	logger "github.com/Yu-Jack/shop-ddd-go/kit/logger"
	"github.com/gin-gonic/gin"
)

type ctrl struct {
	r         *gin.Engine
	orderUc   orderUc.Usecase
	orderSaga orderUc.Saga
}

type Ctrl interface {
	Route()
}

func New(r *gin.Engine, orderUc orderUc.Usecase, orderSaga orderUc.Saga) Ctrl {
	return &ctrl{
		r:         r,
		orderUc:   orderUc,
		orderSaga: orderSaga,
	}
}

func (ct *ctrl) Route() {
	ct.r.POST("/order/checkout", ct.checkoutOrder)
	ct.r.GET("/order/:id", ct.getOrder)
	ct.r.GET("/orders", ct.getOrders)
}

func (ct *ctrl) checkoutOrder(ctx *gin.Context) {
	var req CreateOrderReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, error_code.New(error_code.REQUEST_BODY_IS_INVALID))
		return
	}

	o, err := ct.orderUc.CheckoutOrder(ctx, orderUc.CheckoutOrderInput{
		ConsumerID: req.ConsumerID,
	}, ct.orderSaga)

	if err != nil {
		logger.Log(ctx, "err", err)
		ctx.JSON(200, error_code.New(error_code.ORDER_IS_NOT_AVAILABLE))
		return
	}

	ctx.JSON(200, o)
}

func (ct *ctrl) getOrder(c *gin.Context) {
	var req GetOrderReq
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, error_code.New(error_code.REQUEST_URI_IS_INVALID))
		return
	}

	o, _ := ct.orderUc.FindOrderById(c, req.ID)

	c.JSON(200, o)
}

func (ct *ctrl) getOrders(c *gin.Context) {
	orders, _ := ct.orderUc.GetAllOrders(c)
	c.JSON(200, gin.H{
		"orders": orders,
	})
}
