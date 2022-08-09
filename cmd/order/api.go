package order

import (
	repository "github.com/Yu-Jack/shop-ddd-go/internal/adapter/repository/mysql/order"
	router "github.com/Yu-Jack/shop-ddd-go/internal/router/handler/order"
	usecase "github.com/Yu-Jack/shop-ddd-go/internal/usecase/order"
	"github.com/Yu-Jack/shop-ddd-go/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	r := gin.Default()
	r.Use(logger.Middleware())

	repo := repository.New(db)
	orderUsecase := usecase.NewOrder(repo)
	orderItemUsecase := usecase.NewOrderItem(repo)
	router.New(r, orderUsecase, orderItemUsecase).Route()

	r.Run()
}
