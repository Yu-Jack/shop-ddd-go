package order

import (
	usecase "github.com/Yu-Jack/shop-ddd-go/internal/usecase/order"
	"github.com/Yu-Jack/shop-ddd-go/pkg/logger"
	"github.com/gin-gonic/gin"
)

type route struct {
	engine *gin.Engine

	orderUc     usecase.Order
	orderItemUc usecase.OrderItem
}

type Router interface {
	Route()
}

func New(engine *gin.Engine, orderUc usecase.Order, orderItemUc usecase.OrderItem) Router {
	return &route{
		engine:      engine,
		orderUc:     orderUc,
		orderItemUc: orderItemUc,
	}
}

func (r *route) Route() {
	r.engine.POST("/order/checkout", r.checkoutOrder)
	r.engine.GET("/order/:id", r.getOrder)
	r.engine.GET("/orders", r.getOrders)
	r.engine.POST("/order/item", r.createOrderItem)
	r.engine.GET("/order/item/:order_id", r.getOrderItems)
}

func (r *route) checkoutOrder(ctx *gin.Context) {
	var req CreateOrderReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, nil)
		return
	}

	o, err := r.orderUc.CheckoutOrder(ctx, usecase.CheckoutOrderInput{
		ConsumerID: req.ConsumerID,
	})

	if err != nil {
		logger.Log(ctx, "err", err)
		ctx.JSON(200, nil)
		return
	}

	ctx.JSON(200, o)
}

func (r *route) getOrder(c *gin.Context) {
	var req GetOrderReq
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, nil)
		return
	}

	o, _ := r.orderUc.FindOrderById(c, req.ID)

	c.JSON(200, o)
}

func (r *route) getOrders(c *gin.Context) {
	orders, _ := r.orderUc.GetAllOrders(c)
	c.JSON(200, gin.H{
		"orders": orders,
	})
}

func (r *route) createOrderItem(c *gin.Context) {
	var req CreateOrderItemReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	o, _ := r.orderUc.FindAvailableOrderByConsumerId(c, req.ConsumerID)
	oi, _ := r.orderItemUc.CreateOrderItem(c, usecase.CreateOrderItemInput{
		OrderID:    o.ID,
		Name:       req.Name,
		Amount:     req.Amount,
		ConsumerID: req.ConsumerID,
	})
	c.JSON(200, oi)
}

func (r *route) getOrderItems(c *gin.Context) {
	var req GetOrderItemReq
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	ois, _ := r.orderItemUc.GetOrderItems(c, req.OrderID)
	c.JSON(200, ois)
}
