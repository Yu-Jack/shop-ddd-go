package main

import (
	shopRepo "github.com/Yu-Jack/shop-ddd-go/internal/adapter/repository/mysql/shop"
	shopRoute "github.com/Yu-Jack/shop-ddd-go/internal/router/handler/shop"
	shopUsecase "github.com/Yu-Jack/shop-ddd-go/internal/usecase/shop"
	"github.com/Yu-Jack/shop-ddd-go/pkg/dddcore"
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

	eventBus := dddcore.NewEventBus()
	repo := shopRepo.New(db)
	orderUsecase := shopUsecase.NewOrder(repo, eventBus)
	orderItemUsecase := shopUsecase.NewOrderItem(repo)
	shopRoute.New(r, orderUsecase, orderItemUsecase).Route()

	r.Run()
}
