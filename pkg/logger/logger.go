package logger

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/cnrywjd11/go-echo-pkg-oriented-layout/pkg/requestid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

const (
	// Call stack dump size
	StackDumpSize   = 1024 * 4
	DefaultLogLevel = "debug"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetLevel(LogLevel(DefaultLogLevel))
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		FieldMap:        logrus.FieldMap{"msg": "message"},
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	})
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

// Create panic message by combining current goroutine's call stack with error message
func GetPanicTrace(err interface{}) string {
	trace := make([]byte, StackDumpSize)
	traceLength := runtime.Stack(trace, false)
	return fmt.Sprintf("%v\n%s", err, trace[:traceLength])
}

func LogLevel(level string) logrus.Level {
	lv, err := logrus.ParseLevel(level)
	if err != nil {
		panic(err)
	}
	return lv
}

func JsonLogger() *logrus.Logger {
	return logger
}

func JsonLoggerWithRequestID(ctx context.Context) *logrus.Entry {
	requestID := requestid.FromContext(ctx)
	return logger.WithField("requestId", requestID)
}

func SetLogLevel(level string) {
	logger.SetLevel(LogLevel(level))
}

func WithContext(ctx context.Context) *logrus.Logger {
	return logger.WithContext(ctx).Logger
}

func MiddlewareLogSkipper() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			if strings.Contains(c.Request().RequestURI, "liveness") ||
				strings.Contains(c.Request().RequestURI, "readiness") {
				return true
			}
			return false
		},
	})
}
