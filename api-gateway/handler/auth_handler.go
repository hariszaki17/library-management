package handler

import (
	"context"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/hariszaki17/library-management/api-gateway/handler/dto"
	"github.com/hariszaki17/library-management/api-gateway/handler/middleware"
	"github.com/hariszaki17/library-management/api-gateway/logging"
	pb "github.com/hariszaki17/library-management/proto/gen/user/proto"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var jwtSecret = []byte("your-secret-key")

type AuthHandler struct {
	userRPC pb.UserServiceClient
}

func NewAuthHandler(g *echo.Group, userRPC pb.UserServiceClient) {
	handler := &AuthHandler{
		userRPC: userRPC,
	}
	g.POST("/login", handler.Login)
}

// @Summary Login
// @Description Authenticates a user and returns a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param login body dto.AuthRequest true "Login Request"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Invalid request"
// @Failure 401 {string} string "Invalid username or password"
// @Failure 500 {string} string "Internal server error"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c echo.Context) error {
	requestID := middleware.GetRequestID(c)

	var req dto.AuthRequest
	if err := c.Bind(&req); err != nil {
		logrus.WithError(err).Error("Failed to bind login request")
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	logging.Logger.WithFields(logrus.Fields{
		"handler":   "GetUser",
		"username":  req.Username,
		"requestID": requestID,
	}).Info("Fetching user details")

	ctx := context.WithValue(c.Request().Context(), middleware.RequestIDKey, requestID)

	user, err := h.userRPC.Authenticate(ctx, &pb.AuthenticateRequest{Username: req.Username, Password: req.Password})
	if err != nil {
		logrus.WithError(err).Error("Authentication failed")
		return c.JSON(http.StatusUnauthorized, "Invalid username or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.User.Id,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		logrus.WithError(err).Error("Failed to generate JWT token")
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	logrus.WithFields(logrus.Fields{
		"user_id":  user.User.Id,
		"username": user.User.Username,
	}).Info("User authenticated successfully")

	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}
