package user

import (
	"github.com/gin-gonic/gin"

	userDomain "github.com/Yu-Jack/shop-ddd-go/internal/domain/user"
	userUC "github.com/Yu-Jack/shop-ddd-go/internal/usecase/user"
)

type route struct {
	engine *gin.Engine

	userUC      userUC.User
	userEvent   userUC.UserEvent
	userEventCB userUC.UserEventCacllback
}

type Router interface {
	Route()
	StartEventHandler()
}

func New(engine *gin.Engine, userUC userUC.User) Router {
	return &route{
		engine: engine,
		userUC: userUC,
	}
}

func (r *route) StartEventHandler() {
	r.userEvent.SubscribeOrderCreated(userDomain.ORDER_CREATED_EVENT, r.userEventCB.OrderCreatedCallback)
}

func (r *route) Route() {
	r.engine.POST("/user", r.CreateUser)
	r.engine.GET("/users", r.GetUsers)
}

func (r *route) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	user := r.userUC.CreateUser(userUC.CreateUserInput{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Amount:    req.Amount,
	})
	c.JSON(200, user)
}

func (r *route) GetUsers(c *gin.Context) {
	users := r.userUC.GetAllUsers()
	c.JSON(200, gin.H{
		"users": users,
	})
}
