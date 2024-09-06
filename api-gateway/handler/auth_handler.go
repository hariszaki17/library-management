package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/hariszaki17/library-management/api-gateway/config"
	"github.com/hariszaki17/library-management/api-gateway/handler/dto"
	"github.com/hariszaki17/library-management/api-gateway/handler/middleware"
	"github.com/hariszaki17/library-management/proto/constants"
	pb "github.com/hariszaki17/library-management/proto/gen/user/proto"
	"github.com/hariszaki17/library-management/proto/logging"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

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
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body dto.AuthRequest true "Login Request"
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {string} string "Invalid request"
// @Failure 401 {string} string "Invalid username or password"
// @Failure 500 {string} string "Internal server error"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c echo.Context) error {
	requestID := middleware.GetRequestID(c)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "Login",
		"requestID":   requestID,
	}).Info("User try to login")

	var req dto.AuthRequest
	if err := c.Bind(&req); err != nil {
		logger.WithError(err).Error("Failed to bind login request")
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid request"})
	}

	// Validate the request
	var validate = validator.New()
	if err := validate.Struct(req); err != nil {
		logrus.WithError(err).Error("Validation failed")
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid request"})
	}

	// Create gRPC context with metadata
	md := metadata.Pairs(
		constants.RequestIDKeyCtx, requestID)
	grpcCtx := metadata.NewOutgoingContext(c.Request().Context(), md)

	res, err := h.userRPC.Authenticate(grpcCtx, &pb.AuthenticateRequest{Username: req.Username, Password: req.Password})
	if err != nil {
		logger.WithError(err).Error("Authentication failed")
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "Invalid username or password"})
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "Login",
		"requestID":   requestID,
	}).Info("User login authenticated successfully")
	return c.JSON(http.StatusOK, dto.ToAuthResponse(res.Token))
}
