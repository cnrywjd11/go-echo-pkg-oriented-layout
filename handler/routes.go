package handler

import (
	"net/http"

	v1 "github.com/cnrywjd11/go-echo-pkg-oriented-layout/handler/v1"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/liveness", RespondNoContentHandler)
	e.GET("/readiness", RespondNoContentHandler)

	e.GET("/api/v1/hello", v1.HelloHandler)
}

func RespondNoContentHandler(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
