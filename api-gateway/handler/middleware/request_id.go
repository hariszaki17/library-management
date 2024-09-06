package middleware

import (
	"github.com/hariszaki17/library-management/proto/constants"
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
)

func RequestIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestID := xid.New().String()
		c.Set(constants.RequestIDKeyCtx, requestID)
		c.Response().Header().Set(echo.HeaderXRequestID, requestID)
		return next(c)
	}
}

func GetRequestID(c echo.Context) string {
	if reqID := c.Get(constants.RequestIDKeyCtx); reqID != nil {
		if requestID, ok := reqID.(string); ok {
			return requestID
		}
	}
	return ""
}

