package network

import (
	consumerUc "github.com/Yu-Jack/shop-ddd-go-consumer/internal/consumer/usecase"
	"github.com/gin-gonic/gin"
)

type net struct {
	r          *gin.Engine
	consumerUc consumerUc.Usecase
}

type Net interface {
	Route()
}

func New(r *gin.Engine, consumerUc consumerUc.Usecase) Net {
	return &net{
		r:          r,
		consumerUc: consumerUc,
	}
}

func (n *net) Route() {
	n.r.POST("/consumer", n.CreateConsumer)
	n.r.GET("/consumers", n.GetConsumers)
}

func (n *net) CreateConsumer(c *gin.Context) {
	var req CreateConsumerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	consumer := n.consumerUc.CreateConsumer(consumerUc.CreateConsumerInput{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Amount:    req.Amount,
	})
	c.JSON(200, consumer)
}

func (n *net) GetConsumers(c *gin.Context) {
	consumers := n.consumerUc.GetAllConsumers()
	c.JSON(200, gin.H{
		"consumers": consumers,
	})
}
