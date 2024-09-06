// handler/grpc_handler.go
package handler

import (
	"context"
	"errors"
	"time"

	"github.com/hariszaki17/library-management/author-service/config"
	"github.com/hariszaki17/library-management/author-service/handler/dto"
	"github.com/hariszaki17/library-management/author-service/models"
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
	authors, err := r.authorUsecase.GetAuthors(ctx, int(req.Page), int(req.Limit))
	if err != nil {
		logger.WithError(err).Error("Error while calling method authorUsecase.GetAuthors")
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "GetAuthors",
		"userID":      userID,
	}).Info("Finished RPC - GetAuthors")
	return dto.ToGetAuthorsResponse(authors), nil
}

func (r *rpc) CreateAuthor(ctx context.Context, req *pb.CreateAuthorRequest) (*pb.CreateAuthorResponse, error) {
	requestID := utils.ExtractRequestID(ctx)
	userID := utils.ExtractUserID(ctx)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "CreateAuthor",
		"userID":      userID,
	}).Info("Invoke RPC - CreateAuthor")
	ctx = context.WithValue(ctx, constants.RequestIDKeyCtx, requestID)

	// Parse the string into a time.Time object
	parsedDate, err := time.Parse(constants.FormatDate, req.Birthdate)
	if err != nil {
		logger.WithError(err).Error("Error parse Birthdate")
		return nil, err
	}
	_, err = r.authorUsecase.CreateAuthor(ctx, &models.Author{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Biography: req.Biography,
		BirthDate: parsedDate,
	})
	if err != nil {
		logger.WithError(err).Error("Error while calling method authorUsecase.CreateAuthor")
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "CreateAuthor",
		"userID":      userID,
	}).Info("Finished RPC - CreateAuthor")
	return dto.ToCreateAuthorResponse("successfully create an author"), nil
}

func (r *rpc) UpdateAuthor(ctx context.Context, req *pb.UpdateAuthorRequest) (*pb.UpdateAuthorResponse, error) {
	requestID := utils.ExtractRequestID(ctx)
	userID := utils.ExtractUserID(ctx)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "UpdateAuthor",
		"userID":      userID,
	}).Info("Invoke RPC - UpdateAuthor")
	ctx = context.WithValue(ctx, constants.RequestIDKeyCtx, requestID)

	if req.Id < 1 {
		logger.Error("Error while calling method authorUsecase.UpdateAuthor, id must be > 0")
		return nil, errors.New("id must be > 0")
	}
	data := req.GetData().AsMap() // Convert Struct to map[string]interface{}

	_, err := r.authorUsecase.UpdateAuthor(ctx, uint(req.Id), data)
	if err != nil {
		logger.WithError(err).Error("Error while calling method authorUsecase.UpdateAuthor")
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "UpdateAuthor",
		"userID":      userID,
	}).Info("Finished RPC - UpdateAuthor")
	return dto.ToUpdateAuthorResponse("successfully update an author"), nil
}

func (r *rpc) DeleteAuthor(ctx context.Context, req *pb.DeleteAuthorRequest) (*pb.DeleteAuthorResponse, error) {
	requestID := utils.ExtractRequestID(ctx)
	userID := utils.ExtractUserID(ctx)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "DeleteAuthor",
		"userID":      userID,
	}).Info("Invoke RPC - DeleteAuthor")
	ctx = context.WithValue(ctx, constants.RequestIDKeyCtx, requestID)

	if req.Id < 1 {
		logger.Error("Error while calling method authorUsecase.DeleteAuthor, id must be > 0")
		return nil, errors.New("id must be > 0")
	}

	err := r.authorUsecase.DeleteAuthor(ctx, uint(req.Id))
	if err != nil {
		logger.WithError(err).Error("Error while calling method authorUsecase.DeleteAuthor")
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "DeleteAuthor",
		"userID":      userID,
	}).Info("Finished RPC - DeleteAuthor")
	return dto.ToDeleteAuthorResponse("successfully delete an author"), nil
}
