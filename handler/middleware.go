package handler

import (
	"github.com/cnrywjd11/go-echo-pkg-oriented-layout/pkg/logger"
	"github.com/cnrywjd11/go-echo-pkg-oriented-layout/pkg/requestid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterMiddleware(e *echo.Echo) {
	e.Use(
		middleware.Recover(),
		logger.MiddlewareLogSkipper(),
		requestid.MiddlewareRequestID(),
	)
}
