package logger

import (
	"context"
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	defaultLogger log.Logger
)

// InitLogger 初始化日志
func InitLogger(serviceName string) {
	logger := log.With(
		log.NewStdLogger(os.Stdout),
		"service", serviceName,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)
	defaultLogger = logger
	log.SetLogger(logger)
}

// GetLogger 获取日志实例
func GetLogger() log.Logger {
	if defaultLogger == nil {
		InitLogger("default")
	}
	return defaultLogger
}

// Helper 辅助函数
type Helper struct {
	logger log.Logger
}

// NewHelper 创建日志辅助器
func NewHelper(logger log.Logger) *Helper {
	return &Helper{logger: logger}
}

func (h *Helper) Info(ctx context.Context, format string, args ...interface{}) {
	_ = log.WithContext(ctx, h.logger).Log(log.LevelInfo, "msg", fmt.Sprintf(format, args...))
}

func (h *Helper) Warn(ctx context.Context, format string, args ...interface{}) {
	_ = log.WithContext(ctx, h.logger).Log(log.LevelWarn, "msg", fmt.Sprintf(format, args...))
}

func (h *Helper) Error(ctx context.Context, format string, args ...interface{}) {
	_ = log.WithContext(ctx, h.logger).Log(log.LevelError, "msg", fmt.Sprintf(format, args...))
}

func (h *Helper) Debug(ctx context.Context, format string, args ...interface{}) {
	_ = log.WithContext(ctx, h.logger).Log(log.LevelDebug, "msg", fmt.Sprintf(format, args...))
}


