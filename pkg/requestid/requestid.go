package requestid

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const requestIDHeaderKey = "X-Request-ID"
const requestIDKey = "requestID"

func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}

func FromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if reqID, ok := ctx.Value(requestIDKey).(string); ok {
		return reqID
	}
	return ""
}

func MiddlewareRequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestID := c.Request().Header.Get(requestIDHeaderKey)
			if requestID == "" {
				requestID = uuid.New().String()
				c.Request().Header.Set(requestIDHeaderKey, requestID)
			}
			c.Response().Header().Set(requestIDHeaderKey, requestID)

			ctx := WithRequestID(c.Request().Context(), requestID)
			c.SetRequest(c.Request().WithContext(ctx))

			c.Set(requestIDKey, requestID)

			return next(c)
		}
	}
}
