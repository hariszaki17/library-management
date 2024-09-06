// handler/grpc_handler.go
package handler

import (
	"context"
	"errors"

	"github.com/hariszaki17/library-management/category-service/config"
	"github.com/hariszaki17/library-management/category-service/handler/dto"
	"github.com/hariszaki17/library-management/category-service/models"
	"github.com/hariszaki17/library-management/category-service/usecase"
	"github.com/hariszaki17/library-management/proto/constants"
	pb "github.com/hariszaki17/library-management/proto/gen/category/proto"
	"github.com/hariszaki17/library-management/proto/logging"
	"github.com/hariszaki17/library-management/proto/utils"
	"github.com/sirupsen/logrus"
)

func NewRPC(useCase usecase.CategoryUsecase) pb.CategoryServiceServer {
	return &rpc{
		categoryUsecase: useCase,
	}
}

type rpc struct {
	categoryUsecase                       usecase.CategoryUsecase
	pb.UnimplementedCategoryServiceServer // Embed the unimplemented server
}

func (r *rpc) GetCategories(ctx context.Context, req *pb.GetCategoriesRequest) (*pb.GetCategoriesResponse, error) {
	requestID := utils.ExtractRequestID(ctx)
	userID := utils.ExtractUserID(ctx)
	logger := logging.Logger.WithField("requestID", requestID)
	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "GetCategories",
		"userID":      userID,
	}).Info("Invoke RPC - GetCategories")
	ctx = context.WithValue(ctx, constants.RequestIDKeyCtx, requestID)
	categories, err := r.categoryUsecase.GetCategories(ctx, int(req.Page), int(req.Limit))
	if err != nil {
		logger.WithError(err).Error("Error while calling method categoryUsecase.GetCategories")
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "GetCategories",
		"userID":      userID,
	}).Info("Finished RPC - GetCategories")
	return dto.ToGetCategoriesResponse(categories), nil
}

func (r *rpc) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {
	requestID := utils.ExtractRequestID(ctx)
	userID := utils.ExtractUserID(ctx)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "CreateCategory",
		"userID":      userID,
	}).Info("Invoke RPC - CreateCategory")
	ctx = context.WithValue(ctx, constants.RequestIDKeyCtx, requestID)
	_, err := r.categoryUsecase.CreateCategory(ctx, &models.Category{
		Name: req.Name,
	})
	if err != nil {
		logger.WithError(err).Error("Error while calling method categoryUsecase.CreateCategory")
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "CreateCategory",
		"userID":      userID,
	}).Info("Finished RPC - CreateCategory")
	return dto.ToCreateCategoryResponse("successfully create a book"), nil
}

func (r *rpc) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.UpdateCategoryResponse, error) {
	requestID := utils.ExtractRequestID(ctx)
	userID := utils.ExtractUserID(ctx)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "UpdateCategory",
		"userID":      userID,
	}).Info("Invoke RPC - UpdateCategory")
	ctx = context.WithValue(ctx, constants.RequestIDKeyCtx, requestID)

	if req.Id < 1 {
		logger.Error("Error while calling method categoryUsecase.UpdateCategory, id must be > 0")
		return nil, errors.New("id must be > 0")
	}
	data := req.GetData().AsMap() // Convert Struct to map[string]interface{}

	_, err := r.categoryUsecase.UpdateCategory(ctx, uint(req.Id), data)
	if err != nil {
		logger.WithError(err).Error("Error while calling method categoryUsecase.UpdateCategory")
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "UpdateCategory",
		"userID":      userID,
	}).Info("Finished RPC - UpdateCategory")
	return dto.ToUpdateCategoryResponse("successfully update a book"), nil
}

func (r *rpc) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryResponse, error) {
	requestID := utils.ExtractRequestID(ctx)
	userID := utils.ExtractUserID(ctx)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "DeleteCategory",
		"userID":      userID,
	}).Info("Invoke RPC - DeleteCategory")
	ctx = context.WithValue(ctx, constants.RequestIDKeyCtx, requestID)

	if req.Id < 1 {
		logger.Error("Error while calling method categoryUsecase.DeleteCategory, id must be > 0")
		return nil, errors.New("id must be > 0")
	}

	err := r.categoryUsecase.DeleteCategory(ctx, uint(req.Id))
	if err != nil {
		logger.WithError(err).Error("Error while calling method categoryUsecase.DeleteCategory")
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "DeleteCategory",
		"userID":      userID,
	}).Info("Finished RPC - DeleteCategory")
	return dto.ToDeleteCategoryResponse("successfully delete a book"), nil
}
