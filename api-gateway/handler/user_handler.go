package handler

import (
	"net/http"
	"strconv"

	"github.com/hariszaki17/library-management/api-gateway/config"
	"github.com/hariszaki17/library-management/api-gateway/handler/dto"
	"github.com/hariszaki17/library-management/api-gateway/handler/middleware"
	"github.com/hariszaki17/library-management/proto/constants"
	pb "github.com/hariszaki17/library-management/proto/gen/user/proto"
	"github.com/hariszaki17/library-management/proto/logging"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/metadata"

	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	userRPC pb.UserServiceClient
}

func NewUserHandler(g *echo.Group, userRPC pb.UserServiceClient, authMiddleware echo.MiddlewareFunc) {
	handler := &UserHandler{
		userRPC: userRPC,
	}

	g.GET("/:id", handler.GetUser, authMiddleware)
}

// / GetUser godoc
// @Summary Get a user by ID
// @Description Get a user by their ID from the gRPC service
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param Authorization header string true "Bearer token" Example: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjU4NTMyODksImlkIjoyLCJ1c2VybmFtZSI6ImxhbGFsYSJ9.GYCeJu8qbggiWD14wzwwvXxn6VeVczFobXZzXl-8H6Y"
// @Success 200 {object} dto.GetUserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c echo.Context) error {
	strId := c.Param("id")
	requestID := middleware.GetRequestID(c)
	userID := middleware.GetUserID(c)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "GetUser",
		"userID":      userID,
	}).Info("Fetching user details")

	id, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		logger.WithError(err).Error("Invalid user ID")
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid user ID"})
	}

	// Create gRPC context with metadata
	md := metadata.Pairs(
		constants.RequestIDKeyCtx, requestID,
		constants.UserIDKeyCtx, userID)
	grpcCtx := metadata.NewOutgoingContext(c.Request().Context(), md)

	req := &pb.GetUserDetailsRequest{Id: id}
	resp, err := h.userRPC.GetUserDetails(grpcCtx, req)
	if err != nil {
		logger.WithError(err).Error("Failed to get user details from gRPC server")
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Internal server error"})
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "GetUser",
		"userID":      userID,
		"username":    resp.User.Username,
	}).Info("Successfully fetched user details")

	return c.JSON(http.StatusOK, dto.ToGetUserResponse(resp.User))
}
