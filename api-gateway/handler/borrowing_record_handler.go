package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/hariszaki17/library-management/api-gateway/config"
	"github.com/hariszaki17/library-management/api-gateway/handler/dto"
	"github.com/hariszaki17/library-management/api-gateway/handler/middleware"
	"github.com/hariszaki17/library-management/api-gateway/helper"
	"github.com/hariszaki17/library-management/proto/constants"
	pb "github.com/hariszaki17/library-management/proto/gen/user/proto"
	"github.com/hariszaki17/library-management/proto/logging"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/metadata"

	"github.com/sirupsen/logrus"
)

type BorrowingRecordHandler struct {
	userRPC pb.UserServiceClient
}

func NewBorrowingRecordHandler(g *echo.Group, userRPC pb.UserServiceClient, authMiddleware echo.MiddlewareFunc) {
	handler := &BorrowingRecordHandler{
		userRPC: userRPC,
	}

	g.POST("", handler.UserBorrowBook, authMiddleware)
	g.GET("/records", handler.GetBorrowingRecords, authMiddleware)
	g.POST("/return", handler.UserReturnBook, authMiddleware)
}

// BorrowBook godoc
// @Summary Create Borrowing Record
// @Description User borrow a book from the gRPC service
// @Tags BorrowBook
// @Accept json
// @Produce json
// @Param book body dto.UserBorrowBookRequest true "User borrow book information"
// @Param Authorization header string true "Bearer token" Example: Bearer xxx"
// @Success 200 {object} dto.UserBorrowBookResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /borrow-book [post]
func (h *BorrowingRecordHandler) UserBorrowBook(c echo.Context) error {
	requestID := middleware.GetRequestID(c)
	userID := middleware.GetUserID(c)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "UserBorrowBook",
		"userID":      userID,
	}).Info("User try to borrow book")

	var req *dto.UserBorrowBookRequest
	if err := c.Bind(&req); err != nil {
		logger.WithError(err).Error("Failed to bind user borrow book request")
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid request"})
	}

	// Validate the request
	var validate = validator.New()
	if err := validate.Struct(req); err != nil {
		logrus.WithError(err).Error("Validation failed")
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid request"})
	}

	userIDUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		logger.WithError(err).Error("Invalid user ID")
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid user ID"})
	}

	// Create gRPC context with metadata
	md := metadata.Pairs(
		constants.RequestIDKeyCtx, requestID,
		constants.UserIDKeyCtx, userID)
	grpcCtx := metadata.NewOutgoingContext(c.Request().Context(), md)

	reqRPC := &pb.UserBorrowBookRequest{BookId: uint64(req.BookID), UserId: userIDUint}
	_, err = h.userRPC.UserBorrowBook(grpcCtx, reqRPC)
	if err != nil {
		message := "Internal server error"
		if strings.Contains(err.Error(), "out of stock") {
			message = "Book out of stock"
		}
		logger.WithError(err).Error("Failed to user borrow book from gRPC server")
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: message})
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "UserBorrowBook",
		"userID":      userID,
	}).Info("Successfully user borrow a book")

	return c.JSON(http.StatusOK, dto.ToUserBorrowBookResponse("Successfully user borrow a book"))
}

// ReturnBook godoc
// @Summary Update borrowing record and return the book
// @Description User return a book from the gRPC service
// @Tags BorrowBook
// @Accept json
// @Produce json
// @Param book body dto.UserReturnBookRequest true "User return book information"
// @Param Authorization header string true "Bearer token" Example: Bearer xxx"
// @Success 200 {object} dto.UserReturnBookResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /borrow-book/return [post]
func (h *BorrowingRecordHandler) UserReturnBook(c echo.Context) error {
	requestID := middleware.GetRequestID(c)
	userID := middleware.GetUserID(c)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "UserReturnBook",
		"userID":      userID,
	}).Info("User try to return book")

	var req *dto.UserReturnBookRequest
	if err := c.Bind(&req); err != nil {
		logger.WithError(err).Error("Failed to bind user return book request")
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
		constants.RequestIDKeyCtx, requestID,
		constants.UserIDKeyCtx, userID)
	grpcCtx := metadata.NewOutgoingContext(c.Request().Context(), md)

	reqRPC := &pb.UserReturnBookRequest{Id: uint64(req.ID)}
	_, err := h.userRPC.UserReturnBook(grpcCtx, reqRPC)
	if err != nil {
		logger.WithError(err).Error("Failed to user return book from gRPC server")
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Internal server error"})
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "UserReturnBook",
		"userID":      userID,
	}).Info("Successfully user return a book")

	return c.JSON(http.StatusOK, dto.ToUserReturnBookResponse("Successfully user return a book"))
}

// GetBorrowingRecords godoc
// @Summary Get a list of borrowing records
// @Description Retrieve a paginated list of borrowing records from the gRPC service
// @Tags BorrowBook
// @Accept json
// @Produce json
// @Param page query int true "Page number" default(1)
// @Param limit query int true "Number of items per page" default(10)
// @Param Authorization header string true "Bearer token" Example: Bearer xxx"
// @Success 200 {object} dto.GetBorrowingRecordsResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /borrow-book/records [get]
func (h *BorrowingRecordHandler) GetBorrowingRecords(c echo.Context) error {
	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")

	requestID := middleware.GetRequestID(c)
	userID := middleware.GetUserID(c)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "GetBorrowingRecords",
		"userID":      userID,
	}).Info("Fetching borrowing records")

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

	req := &pb.GetBorrowingRecordsRequest{Page: uint64(page), Limit: uint64(limit)}
	resp, err := h.userRPC.GetBorrowingRecords(grpcCtx, req)
	if err != nil {
		logger.WithError(err).Error("Failed to get books from gRPC server")
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Internal server error"})
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "GetBorrowingRecords",
		"userID":      userID,
	}).Info("Successfully fetched borrowing books")

	return c.JSON(http.StatusOK, dto.ToGetBorrowingRecordsResponse(resp.BorrowingRecords))
}
