package logger

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	log "github.com/go-kit/log"
	"github.com/google/uuid"
)

const LOGGER_NAME = "logger"

func defaultLogger() (logger log.Logger) {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "request_id", uuid.NewString())
	logger.Log()
	return logger
}

func Middleware() func(ctx *gin.Context) {
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

func Log(ctx context.Context, keyvals ...interface{}) {
	_, file, no, _ := runtime.Caller(1)
	full := fmt.Sprintf("%s:%d", file, no)

	logger := ctx.Value(LOGGER_NAME).(log.Logger)
	logger = log.With(logger, "caller", full)

	logger.Log(keyvals...)
}
