// handler/grpc_handler.go
package handler

import (
	"context"

	"github.com/hariszaki17/library-management/category-service/config"
	"github.com/hariszaki17/library-management/category-service/handler/dto"
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
	books, err := r.categoryUsecase.GetCategories(ctx, int(req.Page), int(req.Limit))
	if err != nil {
		logger.WithError(err).Error("Error while calling method categoryUsecase.GetCategories")
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "GetCategories",
		"userID":      userID,
	}).Info("Finished RPC - GetCategories")
	return dto.ToGetCategoriesResponse(books), nil
}
