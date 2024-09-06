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
	pb "github.com/hariszaki17/library-management/proto/gen/category/proto"
	"github.com/hariszaki17/library-management/proto/logging"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/metadata"

	"github.com/sirupsen/logrus"
)

type CategoryHandler struct {
	categoryRPC pb.CategoryServiceClient
}

func NewCategoryHandler(g *echo.Group, categoryRPC pb.CategoryServiceClient, authMiddleware echo.MiddlewareFunc) {
	handler := &CategoryHandler{
		categoryRPC: categoryRPC,
	}

	g.GET("", handler.GetCategories, authMiddleware)
	g.POST("", handler.CreateCategory, authMiddleware)
	g.PUT("/:id", handler.UpdateCategory, authMiddleware)
	g.DELETE("/:id", handler.DeleteCategory, authMiddleware)
}

// GetCategories godoc
// @Summary Get a list of categories
// @Description Retrieve a paginated list of categories from the gRPC service
// @Tags Categories
// @Accept json
// @Produce json
// @Param page query int true "Page number" default(1)
// @Param limit query int true "Number of items per page" default(10)
// @Param Authorization header string true "Bearer token" Example: Bearer xxx"
// @Success 200 {object} dto.GetCategoriesResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /categories [get]
func (h *CategoryHandler) GetCategories(c echo.Context) error {
	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")

	requestID := middleware.GetRequestID(c)
	userID := middleware.GetUserID(c)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "GetCategories",
		"userID":      userID,
	}).Info("Fetching categories")

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

	req := &pb.GetCategoriesRequest{Page: uint64(page), Limit: uint64(limit)}
	resp, err := h.categoryRPC.GetCategories(grpcCtx, req)
	if err != nil {
		logger.WithError(err).Error("Failed to get categories from gRPC server")
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Internal server error"})
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "GetCategories",
		"userID":      userID,
	}).Info("Successfully fetched categories")

	return c.JSON(http.StatusOK, dto.ToGetCategoriesResponse(resp.Categories))
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category with the provided information
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body dto.CreateCategoryRequest true "Category information"
// @Param Authorization header string true "Bearer token" Example: Bearer xxx"
// @Success 200 {object} dto.CreateCategoryResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /categories [post]
func (h *CategoryHandler) CreateCategory(c echo.Context) error {
	requestID := middleware.GetRequestID(c)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "CreateCategory",
		"requestID":   requestID,
	}).Info("User try to create category")

	var req *dto.CreateCategoryRequest
	if err := c.Bind(&req); err != nil {
		logger.WithError(err).Error("Failed to bind create category request")
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

	rpcReq := dto.CreateCategoryRPCRequest(req)
	_, err := h.categoryRPC.CreateCategory(grpcCtx, rpcReq)
	if err != nil {
		logger.WithError(err).Error("Create category failed")
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: fmt.Sprintf("Error while create category with err: %s", err.Error()),
		})
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "Create Category",
		"requestID":   requestID,
	}).Info("Category created successfully")
	return c.JSON(http.StatusOK, dto.ToCreateCategoryResponse("Category created successfully"))
}

// UpdateCategory godoc
// @Summary Update a category
// @Description Update a category with the provided information
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body dto.UpdateCategoryRequest true "Category information"
// @Param Authorization header string true "Bearer token" Example: Bearer xxx"
// @Success 200 {object} dto.UpdateCategoryResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /categories/{id} [put]
func (h *CategoryHandler) UpdateCategory(c echo.Context) error {
	strId := c.Param("id")
	requestID := middleware.GetRequestID(c)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "UpdateCategory",
		"requestID":   requestID,
	}).Info("User try to update category")

	id, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		logger.WithError(err).Error("Invalid category ID")
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid category ID"})
	}

	var req dto.UpdateCategoryRequest
	if err := c.Bind(&req); err != nil {
		logger.WithError(err).Error("Failed to bind update category request")
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

	rpcReq, err := dto.UpdateCategoryRPCRequest(uint(id), req)
	if err != nil {
		logrus.WithError(err).Error("Error while converting request")
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	_, err = h.categoryRPC.UpdateCategory(grpcCtx, rpcReq)
	if err != nil {
		logger.WithError(err).Error("Update category failed")
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: fmt.Sprintf("Error while update category with err: %s", err.Error()),
		})
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "Update Category",
		"requestID":   requestID,
	}).Info("Category update successfully")
	return c.JSON(http.StatusOK, dto.ToUpdateCategoryResponse("Category updated successfully"))
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Delete a category with the provided information
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param Authorization header string true "Bearer token" Example: Bearer xxx"
// @Success 200 {object} dto.DeleteCategoryResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c echo.Context) error {
	strId := c.Param("id")
	requestID := middleware.GetRequestID(c)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "DeleteCategory",
		"requestID":   requestID,
	}).Info("User try to delete category")

	id, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		logger.WithError(err).Error("Invalid category ID")
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid category ID"})
	}

	// Create gRPC context with metadata
	md := metadata.Pairs(
		constants.RequestIDKeyCtx, requestID)
	grpcCtx := metadata.NewOutgoingContext(c.Request().Context(), md)

	rpcReq := dto.DeleteCategoryRPCRequest(uint(id))
	_, err = h.categoryRPC.DeleteCategory(grpcCtx, rpcReq)
	if err != nil {
		logger.WithError(err).Error("Delete category failed")
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: fmt.Sprintf("Error while delete category with err: %s", err.Error()),
		})
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"handler":     "Delete Category",
		"requestID":   requestID,
	}).Info("Category deleted successfully")
	return c.JSON(http.StatusOK, dto.ToDeleteCategoryResponse("Category deleted successfully"))
}
