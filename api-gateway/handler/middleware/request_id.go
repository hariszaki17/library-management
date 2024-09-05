package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
)

const RequestIDKey = "RequestID"

func RequestIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestID := xid.New().String()
		c.Set(RequestIDKey, requestID)
		c.Response().Header().Set(echo.HeaderXRequestID, requestID)
		return next(c)
	}
}

func GetRequestID(c echo.Context) string {
	if reqID := c.Get(RequestIDKey); reqID != nil {
		if requestID, ok := reqID.(string); ok {
			return requestID
		}
	}
	return ""
}
