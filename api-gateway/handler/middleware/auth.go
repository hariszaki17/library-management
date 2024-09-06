package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/hariszaki17/library-management/api-gateway/handler/dto"
	"github.com/hariszaki17/library-management/proto/constants"

	pb "github.com/hariszaki17/library-management/proto/gen/user/proto"
	"github.com/labstack/echo/v4"
)

// AuthJWTMiddleware returns a JWT authentication middleware
func AuthJWTMiddleware(userRPC pb.UserServiceClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Set Content-Type header
			c.Response().Header().Set("Content-Type", "application/json")

			// Get Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "Missing authorization header"})
			}

			// Remove "Bearer " prefix
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "Malformed authorization header"})
			}

			res, err := userRPC.VerifyJWT(c.Request().Context(), &pb.VerifyJWTRequest{
				Token: tokenString,
			})
			if err != nil {
				return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "Failed to verify token"})
			}

			c.Set(constants.UserIDKeyCtx, fmt.Sprintf("%v", res.User.Id))
			return next(c)
		}
	}
}

func GetUserID(c echo.Context) string {
	if userID := c.Get(constants.UserIDKeyCtx); userID != nil {
		if id, ok := userID.(string); ok {
			return id
		}
	}
	return ""
}
