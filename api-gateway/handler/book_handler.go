package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
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
	g.POST("", handler.CreateBook, authMiddleware)
	g.PUT("/:id", handler.UpdateBook, authMiddleware)
	g.DELETE("/:id", handler.DeleteBook, authMiddleware)
	g.GET("/recommendation", handler.GetBookRecommendation, authMiddleware)

}

// GetBooks godoc
// @Summary Get a list of books
// @Description Retrieve a paginated list of books from the gRPC service
// @Tags Books
// @Accept json
// @Produce json
// @Param page query int true "Page number" default(1)
// @Param limit query int true "Number of items per page" default(10)
// @Param query query string false "Query search of title"
// @Param Authorization header string true "Bearer token" Example: Bearer xxx"
// @Success 200 {object} dto.GetBooksResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /books [get]
func (h *BookHandler) GetBooks(c echo.Context) error {
	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")
	query := c.QueryParam("query")

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

	if err := helper.ValidatePageLimit(int(page), int(limit)); err != nil {
		logger.WithError(err).Error("Invalid pagination")
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	// Create gRPC context with metadata
	md := metadata.Pairs(
		constants.RequestIDKeyCtx, requestID,
		constants.UserIDKeyCtx, userID)
	grpcCtx := metadata.NewOutgoingContext(c.Request().Context(), md)

	req := &pb.GetBooksRequest{Page: uint64(page), Limit: uint64(limit), Query: query}
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

// CreateBook godoc
// @Summary Create a new book
// @Description Create a new book with the provided information
// @Tags Books
// @Accept json
// @Produce json
// @Param book body dto.CreateBookRequest true "Book information"
// @Param Authorization header string true "Bearer token" Example: Bearer xxx"
// @Success 200 {object} dto.CreateBookResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /books [post]
func (h *BookHandler) CreateBook(c echo.Context) error {
	requestID := middleware.GetRequestID(c)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "CreateBook",
		"requestID":   requestID,
	}).Info("User try to create book")

	var req *dto.CreateBookRequest
	if err := c.Bind(&req); err != nil {
		logger.WithError(err).Error("Failed to bind create book request")
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

	rpcReq := dto.CreateBookRPCRequest(req)
	_, err := h.bookRPC.CreateBook(grpcCtx, rpcReq)
	if err != nil {
		logger.WithError(err).Error("Create book failed")
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: fmt.Sprintf("Error while create book with err: %s", err.Error()),
		})
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "Create Book",
		"requestID":   requestID,
	}).Info("Book created successfully")
	return c.JSON(http.StatusOK, dto.ToCreateBookResponse("Book created successfully"))
}

// UpdateBook godoc
// @Summary Update a book
// @Description Update a book with the provided information
// @Tags Books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body dto.UpdateBookRequest true "Book information"
// @Param Authorization header string true "Bearer token" Example: Bearer xxx"
// @Success 200 {object} dto.UpdateBookResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /books/{id} [put]
func (h *BookHandler) UpdateBook(c echo.Context) error {
	strId := c.Param("id")
	requestID := middleware.GetRequestID(c)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "UpdateBook",
		"requestID":   requestID,
	}).Info("User try to update book")

	id, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		logger.WithError(err).Error("Invalid book ID")
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid book ID"})
	}

	var req dto.UpdateBookRequest
	if err := c.Bind(&req); err != nil {
		logger.WithError(err).Error("Failed to bind update book request")
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

	rpcReq, err := dto.UpdateBookRPCRequest(uint(id), req)
	if err != nil {
		logrus.WithError(err).Error("Error while converting request")
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	_, err = h.bookRPC.UpdateBook(grpcCtx, rpcReq)
	if err != nil {
		logger.WithError(err).Error("Update book failed")
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: fmt.Sprintf("Error while update book with err: %s", err.Error()),
		})
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "Update Book",
		"requestID":   requestID,
	}).Info("Book update successfully")
	return c.JSON(http.StatusOK, dto.ToUpdateBookResponse("Book updated successfully"))
}

// DeleteBook godoc
// @Summary Delete a book
// @Description Delete a book with the provided information
// @Tags Books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param Authorization header string true "Bearer token" Example: Bearer xxx"
// @Success 200 {object} dto.DeleteBookResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /books/{id} [delete]
func (h *BookHandler) DeleteBook(c echo.Context) error {
	strId := c.Param("id")
	requestID := middleware.GetRequestID(c)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "DeleteBook",
		"requestID":   requestID,
	}).Info("User try to delete book")

	id, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		logger.WithError(err).Error("Invalid book ID")
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid book ID"})
	}

	// Create gRPC context with metadata
	md := metadata.Pairs(
		constants.RequestIDKeyCtx, requestID)
	grpcCtx := metadata.NewOutgoingContext(c.Request().Context(), md)

	rpcReq := dto.DeleteBookRPCRequest(uint(id))
	_, err = h.bookRPC.DeleteBook(grpcCtx, rpcReq)
	if err != nil {
		logger.WithError(err).Error("Delete book failed")
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: fmt.Sprintf("Error while delete book with err: %s", err.Error()),
		})
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "Delete Book",
		"requestID":   requestID,
	}).Info("Book deleted successfully")
	return c.JSON(http.StatusOK, dto.ToDeleteBookResponse("Book deleted successfully"))
}

// GetBookRecommendation godoc
// @Summary Get a list of book recommendation
// @Description Retrieve a list of book recommendation from the gRPC service
// @Tags Books
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token" Example: Bearer xxx"
// @Success 200 {object} dto.GetBookRecommendationResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /books/recommendation [get]
func (h *BookHandler) GetBookRecommendation(c echo.Context) error {
	requestID := middleware.GetRequestID(c)
	userID := middleware.GetUserID(c)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "GetBookRecommendation",
		"userID":      userID,
	}).Info("Fetching book recommendation")

	// Create gRPC context with metadata
	md := metadata.Pairs(
		constants.RequestIDKeyCtx, requestID,
		constants.UserIDKeyCtx, userID)
	grpcCtx := metadata.NewOutgoingContext(c.Request().Context(), md)

	resp, err := h.bookRPC.GetBookRecommendation(grpcCtx, &pb.GetBookRecommendationRequest{})
	if err != nil {
		logger.WithError(err).Error("Failed to get book recommendation from gRPC server")
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Internal server error"})
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "GetBookRecommendation",
		"userID":      userID,
	}).Info("Successfully fetched book book recommendation")

	return c.JSON(http.StatusOK, dto.ToGetBookRecommendationResponse(resp.BookRecommendation))
}
