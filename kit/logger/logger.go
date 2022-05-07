package logger

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/go-kit/kit/log"
	"github.com/google/uuid"
)

var LOGGER_NAME = "logger"

func New() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var logger log.Logger
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
		logger = log.With(logger, "request_id", uuid.NewString())
		ctx.Set(LOGGER_NAME, logger)

		ctx.Next()
	}
}

func GetLogger(ctx context.Context) log.Logger {
	l := ctx.Value(LOGGER_NAME)
	return l.(log.Logger)
}
