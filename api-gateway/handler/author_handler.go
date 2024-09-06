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
	g.POST("", handler.CreateAuthor, authMiddleware)
	g.PUT("/:id", handler.UpdateAuthor, authMiddleware)
	g.DELETE("/:id", handler.DeleteAuthor, authMiddleware)
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
	}).Info("Fetching authors")

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

// CreateAuthor godoc
// @Summary Create a new author
// @Description Create a new author with the provided information
// @Tags Authors
// @Accept json
// @Produce json
// @Param author body dto.CreateAuthorRequest true "Author information"
// @Param Authorization header string true "Bearer token" Example: Bearer xxx"
// @Success 200 {object} dto.CreateAuthorResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /authors [post]
func (h *AuthorHandler) CreateAuthor(c echo.Context) error {
	requestID := middleware.GetRequestID(c)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "CreateAuthor",
		"requestID":   requestID,
	}).Info("User try to create author")

	var req *dto.CreateAuthorRequest
	if err := c.Bind(&req); err != nil {
		logger.WithError(err).Error("Failed to bind create author request")
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

	rpcReq := dto.CreateAuthorRPCRequest(req)
	_, err := h.authorRPC.CreateAuthor(grpcCtx, rpcReq)
	if err != nil {
		logger.WithError(err).Error("Create author failed")
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: fmt.Sprintf("Error while create author with err: %s", err.Error()),
		})
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "Create Author",
		"requestID":   requestID,
	}).Info("Author created successfully")
	return c.JSON(http.StatusOK, dto.ToCreateAuthorResponse("Author created successfully"))
}

// UpdateAuthor godoc
// @Summary Update a author
// @Description Update a author with the provided information
// @Tags Authors
// @Accept json
// @Produce json
// @Param id path int true "Author ID"
// @Param author body dto.UpdateAuthorRequest true "Author information"
// @Param Authorization header string true "Bearer token" Example: Bearer xxx"
// @Success 200 {object} dto.UpdateAuthorResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /authors/{id} [put]
func (h *AuthorHandler) UpdateAuthor(c echo.Context) error {
	strId := c.Param("id")
	requestID := middleware.GetRequestID(c)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "UpdateAuthor",
		"requestID":   requestID,
	}).Info("User try to update author")

	id, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		logger.WithError(err).Error("Invalid author ID")
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid author ID"})
	}

	var req dto.UpdateAuthorRequest
	if err := c.Bind(&req); err != nil {
		logger.WithError(err).Error("Failed to bind update author request")
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

	rpcReq, err := dto.UpdateAuthorRPCRequest(uint(id), req)
	if err != nil {
		logrus.WithError(err).Error("Error while converting request")
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	_, err = h.authorRPC.UpdateAuthor(grpcCtx, rpcReq)
	if err != nil {
		logger.WithError(err).Error("Update author failed")
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: fmt.Sprintf("Error while update author with err: %s", err.Error()),
		})
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "Update Author",
		"requestID":   requestID,
	}).Info("Author update successfully")
	return c.JSON(http.StatusOK, dto.ToUpdateAuthorResponse("Author updated successfully"))
}

// DeleteAuthor godoc
// @Summary Delete a author
// @Description Delete a author with the provided information
// @Tags Authors
// @Accept json
// @Produce json
// @Param id path int true "Author ID"
// @Param Authorization header string true "Bearer token" Example: Bearer xxx"
// @Success 200 {object} dto.DeleteAuthorResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /authors/{id} [delete]
func (h *AuthorHandler) DeleteAuthor(c echo.Context) error {
	strId := c.Param("id")
	requestID := middleware.GetRequestID(c)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "DeleteAuthor",
		"requestID":   requestID,
	}).Info("User try to delete author")

	id, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		logger.WithError(err).Error("Invalid author ID")
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid author ID"})
	}

	// Create gRPC context with metadata
	md := metadata.Pairs(
		constants.RequestIDKeyCtx, requestID)
	grpcCtx := metadata.NewOutgoingContext(c.Request().Context(), md)

	rpcReq := dto.DeleteAuthorRPCRequest(uint(id))
	_, err = h.authorRPC.DeleteAuthor(grpcCtx, rpcReq)
	if err != nil {
		logger.WithError(err).Error("Delete author failed")
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: fmt.Sprintf("Error while delete author with err: %s", err.Error()),
		})
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "Delete Author",
		"requestID":   requestID,
	}).Info("Author deleted successfully")
	return c.JSON(http.StatusOK, dto.ToDeleteAuthorResponse("Author deleted successfully"))
}
