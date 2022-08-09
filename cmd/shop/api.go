package order

import (
	repository "github.com/Yu-Jack/shop-ddd-go/internal/adapter/repository/mysql/shop"
	router "github.com/Yu-Jack/shop-ddd-go/internal/router/handler/shop"
	usecase "github.com/Yu-Jack/shop-ddd-go/internal/usecase/shop"
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
