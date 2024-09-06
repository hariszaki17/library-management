package handler

import (
	"net/http"
	"strconv"

	"github.com/hariszaki17/library-management/api-gateway/config"
	"github.com/hariszaki17/library-management/api-gateway/handler/dto"
	"github.com/hariszaki17/library-management/api-gateway/handler/middleware"
	"github.com/hariszaki17/library-management/api-gateway/helper"
	"github.com/hariszaki17/library-management/proto/constants"
	pb "github.com/hariszaki17/library-management/proto/gen/author/proto"
	"github.com/hariszaki17/library-management/proto/logging"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/metadata"

	"github.com/sirupsen/logrus"
)

type AuthorHandler struct {
	authorRPC pb.AuthorServiceClient
}

func NewAuthorHandler(g *echo.Group, authorRPC pb.AuthorServiceClient, authMiddleware echo.MiddlewareFunc) {
	handler := &AuthorHandler{
		authorRPC: authorRPC,
	}

	g.GET("", handler.GetAuthors, authMiddleware)
}

// GetAuthors godoc
// @Summary Get a list of authors
// @Description Retrieve a paginated list of authors from the gRPC service
// @Tags Authors
// @Accept json
// @Produce json
// @Param page query int true "Page number" default(1)
// @Param limit query int true "Number of items per page" default(10)
// @Param Authorization header string true "Bearer token" Example: Bearer xxx"
// @Success 200 {object} dto.GetAuthorsResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /authors [get]
func (h *AuthorHandler) GetAuthors(c echo.Context) error {
	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")

	requestID := middleware.GetRequestID(c)
	userID := middleware.GetUserID(c)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "GetAuthors",
		"userID":      userID,
	}).Info("Fetching books")

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		logger.WithError(err).Error("Invalid page")
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid page"})
	}

	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		logger.WithError(err).Error("Invalid limit")
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid limit"})
	}

	if err := helper.ValidatePageLimit(int(page), int(limit)); err != nil {
		logger.WithError(err).Error("Invalid pagination")
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	// Create gRPC context with metadata
	md := metadata.Pairs(
		constants.RequestIDKeyCtx, requestID,
		constants.UserIDKeyCtx, userID)
	grpcCtx := metadata.NewOutgoingContext(c.Request().Context(), md)

	req := &pb.GetAuthorsRequest{Page: uint64(page), Limit: uint64(limit)}
	resp, err := h.authorRPC.GetAuthors(grpcCtx, req)
	if err != nil {
		logger.WithError(err).Error("Failed to get authors from gRPC server")
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Internal server error"})
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "GetAuthors",
		"userID":      userID,
	}).Info("Successfully fetched authors")

	return c.JSON(http.StatusOK, dto.ToGetAuthorsResponse(resp.Authors))
}
