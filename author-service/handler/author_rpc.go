// handler/grpc_handler.go
package handler

import (
	"context"

	"github.com/hariszaki17/library-management/author-service/config"
	"github.com/hariszaki17/library-management/author-service/handler/dto"
	"github.com/hariszaki17/library-management/author-service/usecase"
	"github.com/hariszaki17/library-management/proto/constants"
	pb "github.com/hariszaki17/library-management/proto/gen/author/proto"
	"github.com/hariszaki17/library-management/proto/logging"
	"github.com/hariszaki17/library-management/proto/utils"
	"github.com/sirupsen/logrus"
)

func NewRPC(useCase usecase.AuthorUsecase) pb.AuthorServiceServer {
	return &rpc{
		authorUsecase: useCase,
	}
}

type rpc struct {
	authorUsecase                       usecase.AuthorUsecase
	pb.UnimplementedAuthorServiceServer // Embed the unimplemented server
}

func (r *rpc) GetAuthors(ctx context.Context, req *pb.GetAuthorsRequest) (*pb.GetAuthorsResponse, error) {
	requestID := utils.ExtractRequestID(ctx)
	userID := utils.ExtractUserID(ctx)
	logger := logging.Logger.WithField("requestID", requestID)
	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "GetAuthors",
		"userID":      userID,
	}).Info("Invoke RPC - GetAuthors")
	ctx = context.WithValue(ctx, constants.RequestIDKeyCtx, requestID)
	books, err := r.authorUsecase.GetAuthors(ctx, int(req.Page), int(req.Limit))
	if err != nil {
		logger.WithError(err).Error("Error while calling method authorUsecase.GetAuthors")
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "GetAuthors",
		"userID":      userID,
	}).Info("Finished RPC - GetAuthors")
	return dto.ToGetAuthorsResponse(books), nil
}
