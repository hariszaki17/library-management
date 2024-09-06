// handler/grpc_handler.go
package handler

import (
	"context"

	"github.com/hariszaki17/library-management/book-service/config"
	"github.com/hariszaki17/library-management/book-service/handler/dto"
	"github.com/hariszaki17/library-management/book-service/usecase"
	"github.com/hariszaki17/library-management/proto/constants"
	pb "github.com/hariszaki17/library-management/proto/gen/book/proto"
	"github.com/hariszaki17/library-management/proto/logging"
	"github.com/hariszaki17/library-management/proto/utils"
	"github.com/sirupsen/logrus"
)

func NewRPC(useCase usecase.BookUsecase) pb.BookServiceServer {
	return &rpc{
		bookUsecase: useCase,
	}
}

type rpc struct {
	bookUsecase                       usecase.BookUsecase
	pb.UnimplementedBookServiceServer // Embed the unimplemented server
}

func (r *rpc) GetBooks(ctx context.Context, req *pb.GetBooksRequest) (*pb.GetBooksResponse, error) {
	requestID := utils.ExtractRequestID(ctx)
	userID := utils.ExtractUserID(ctx)
	logger := logging.Logger.WithField("requestID", requestID)
	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "GetBooks",
		"userID":      userID,
	}).Info("Invoke RPC - GetBooks")
	ctx = context.WithValue(ctx, constants.RequestIDKeyCtx, requestID)
	books, err := r.bookUsecase.GetBooks(ctx, int(req.Page), int(req.Limit))
	if err != nil {
		logger.WithError(err).Error("Error while calling method bookUsecase.GetBooks")
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "GetBooks",
		"userID":      userID,
	}).Info("Finished RPC - GetBooks")
	return dto.ToGetBooksResponse(books), nil
}
