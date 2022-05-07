package logger

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/go-kit/kit/log"
	"github.com/google/uuid"
)

const LOGGER_NAME = "logger"

func defaultLogger() (logger log.Logger) {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "request_id", uuid.NewString())
	logger = log.With(logger, "caller", log.DefaultCaller)
	return logger
}

func New() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Set(LOGGER_NAME, defaultLogger())
		ctx.Next()
	}
}

func NewContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, LOGGER_NAME, defaultLogger())
	return ctx
}

func GetLogger(ctx context.Context) log.Logger {
	l := ctx.Value(LOGGER_NAME)
	return l.(log.Logger)
}
