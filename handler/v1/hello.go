package v1

import (
	"errors"
	"net/http"

	"github.com/cnrywjd11/go-echo-pkg-oriented-layout/pkg/logger"
	"github.com/labstack/echo/v4"
)

func HelloHandler(c echo.Context) error {
	logger.JsonLogger().Debugf("hello1")
	logger.JsonLogger().Infof("hello2")
	logger.JsonLogger().WithField("key", "value").Infof("hello3")
	logger.JsonLoggerWithRequestID(c.Request().Context()).Infof("hello4")

	// Print panic error
	err := errors.New("hello error")
	logger.JsonLogger().Infof(logger.GetPanicTrace(err))

	return c.String(http.StatusOK, "hello world!")
}
