package handler

import (
	"net/http"
	"strconv"

	"github.com/hariszaki17/library-management/api-gateway/config"
	"github.com/hariszaki17/library-management/api-gateway/handler/dto"
	"github.com/hariszaki17/library-management/api-gateway/handler/middleware"
	"github.com/hariszaki17/library-management/api-gateway/helper"
	"github.com/hariszaki17/library-management/proto/constants"
	pb "github.com/hariszaki17/library-management/proto/gen/book/proto"
	"github.com/hariszaki17/library-management/proto/logging"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/metadata"

	"github.com/sirupsen/logrus"
)

type BookHandler struct {
	bookRPC pb.BookServiceClient
}

func NewBookHandler(g *echo.Group, bookRPC pb.BookServiceClient, authMiddleware echo.MiddlewareFunc) {
	handler := &BookHandler{
		bookRPC: bookRPC,
	}

	g.GET("", handler.GetBooks, authMiddleware)
}

// GetBooks godoc
// @Summary Get a list of books
// @Description Retrieve a paginated list of books from the gRPC service
// @Tags Books
// @Accept json
// @Produce json
// @Param page query int true "Page number" default(1)
// @Param limit query int true "Number of items per page" default(10)
// @Param Authorization header string true "Bearer token" Example: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjU4NTMyODksImlkIjoyLCJ1c2VybmFtZSI6ImxhbGFsYSJ9.GYCeJu8qbggiWD14wzwwvXxn6VeVczFobXZzXl-8H6Y"
// @Success 200 {object} dto.GetBooksResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /books [get]
func (h *BookHandler) GetBooks(c echo.Context) error {
	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")

	requestID := middleware.GetRequestID(c)
	userID := middleware.GetUserID(c)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "GetBooks",
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

	if err := helper.ValidatePageLimit(page, int(limit)); err != nil {
		logger.WithError(err).Error("Invalid pagination")
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	// Create gRPC context with metadata
	md := metadata.Pairs(
		constants.RequestIDKeyCtx, requestID,
		constants.UserIDKeyCtx, userID)
	grpcCtx := metadata.NewOutgoingContext(c.Request().Context(), md)

	req := &pb.GetBooksRequest{Page: uint64(page), Limit: uint64(limit)}
	resp, err := h.bookRPC.GetBooks(grpcCtx, req)
	if err != nil {
		logger.WithError(err).Error("Failed to get books from gRPC server")
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Internal server error"})
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "GetBooks",
		"userID":      userID,
	}).Info("Successfully fetched books")

	return c.JSON(http.StatusOK, dto.ToGetBooksResponse(resp.Books))
}
