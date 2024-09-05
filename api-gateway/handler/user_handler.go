package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/hariszaki17/library-management/api-gateway/handler/dto"
	"github.com/hariszaki17/library-management/api-gateway/handler/middleware"
	"github.com/hariszaki17/library-management/api-gateway/logging"
	pb "github.com/hariszaki17/library-management/proto/gen/user/proto"
	"github.com/labstack/echo/v4"

	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	userRPC pb.UserServiceClient
}

func NewUserHandler(g *echo.Group, userRPC pb.UserServiceClient) {
	handler := &UserHandler{
		userRPC: userRPC,
	}

	g.GET("/:id", handler.GetUser)
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Get a user by their ID from the gRPC service
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dto.GetUserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c echo.Context) error {
	strId := c.Param("id")
	requestID := middleware.GetRequestID(c)
	logging.Logger.WithFields(logrus.Fields{
		"handler":   "GetUser",
		"userID":    strId,
		"requestID": requestID,
	}).Info("Fetching user details")

	id, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		logging.Logger.WithError(err).Error("Invalid user ID")
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid user ID"})
	}

	req := &pb.GetUserDetailsRequest{Id: id}
	ctx := context.WithValue(c.Request().Context(), middleware.RequestIDKey, requestID)
	resp, err := h.userRPC.GetUserDetails(ctx, req)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to get user details from gRPC server")
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Internal server error"})
	}
	logging.Logger.WithFields(logrus.Fields{
		"userID":    id,
		"username":  resp.User.Username,
		"requestID": requestID,
	}).Info("Successfully fetched user details")

	return c.JSON(http.StatusOK, dto.ToGetUserResponse(resp.User))
}
