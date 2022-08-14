package main

import (
	userRepo "github.com/Yu-Jack/shop-ddd-go/internal/adapter/repository/mysql/user"
	userEvent "github.com/Yu-Jack/shop-ddd-go/internal/event/user"
	userRoute "github.com/Yu-Jack/shop-ddd-go/internal/router/handler/user"
	userUc "github.com/Yu-Jack/shop-ddd-go/internal/usecase/user"
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

	userRepository := userRepo.New(db)
	userUsecase := userUc.New(userRepository, eventBus)

	userEvent.New(userUsecase, eventBus).StartEventHandler()
	userRoute.New(r, userUsecase).Route()

	r.Run()
}
