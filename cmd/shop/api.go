package order

import (
	shopRepo "github.com/Yu-Jack/shop-ddd-go/internal/adapter/repository/mysql/shop"
	router "github.com/Yu-Jack/shop-ddd-go/internal/router/handler/shop"
	"github.com/Yu-Jack/shop-ddd-go/internal/usecase/shop/order"
	"github.com/Yu-Jack/shop-ddd-go/internal/usecase/shop/order_item"
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

	repo := shopRepo.NewRepo(db)
	orderUsecase := order.NewUsecase(repo)
	orderItemUsecase := order_item.NewUsecase(repo)
	router.New(r, orderUsecase, orderItemUsecase).Route()

	r.Run()
}
