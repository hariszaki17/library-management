// handler/grpc_handler.go
package handler

import (
	"context"
	"errors"

	"github.com/hariszaki17/library-management/proto/constants"
	pb "github.com/hariszaki17/library-management/proto/gen/user/proto"
	"github.com/hariszaki17/library-management/proto/logging"
	"github.com/hariszaki17/library-management/proto/utils"
	"github.com/hariszaki17/library-management/user-service/config"
	"github.com/hariszaki17/library-management/user-service/handler/dto"
	"github.com/hariszaki17/library-management/user-service/usecase"
	"github.com/sirupsen/logrus"
)

func NewRPC(useCase usecase.UserUsecase, borrowingRecordUsecase usecase.BorrowingRecordUsecase) pb.UserServiceServer {
	return &rpc{
		userUsecase:            useCase,
		borrowingRecordUsecase: borrowingRecordUsecase,
	}
}

type rpc struct {
	userUsecase                       usecase.UserUsecase
	borrowingRecordUsecase            usecase.BorrowingRecordUsecase
	pb.UnimplementedUserServiceServer // Embed the unimplemented server
}

func (r *rpc) GetUserDetails(ctx context.Context, req *pb.GetUserDetailsRequest) (*pb.GetUserDetailsResponse, error) {
	requestID := utils.ExtractRequestID(ctx)
	userID := utils.ExtractUserID(ctx)
	logger := logging.Logger.WithField("requestID", requestID)
	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "GetUserDetails",
		"userID":      userID,
	}).Info("Invoke RPC - GetUserDetails")
	ctx = context.WithValue(ctx, constants.RequestIDKeyCtx, requestID)
	user, err := r.userUsecase.GetUserDetails(ctx, uint(req.Id))
	if err != nil {
		logger.WithError(err).Error("Error while calling method userUsecase.GetUserDetails")
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "GetUserDetails",
		"userID":      userID,
	}).Info("Finished RPC - GetUserDetails")
	return &pb.GetUserDetailsResponse{
		User: &pb.User{
			Username: user.Username,
		},
	}, nil
}

func (r *rpc) Authenticate(ctx context.Context, req *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	requestID := utils.ExtractRequestID(ctx)
	logger := logging.Logger.WithField("requestID", requestID)
	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "Authenticate",
	}).Info("Invoke RPC - Authenticate")

	res, err := r.userUsecase.Authenticate(ctx, req.Username, req.Password)
	if err != nil {
		logger.WithError(err).Error("Error while calling method userUsecase.Authenticate")
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "Authenticate",
	}).Info("Finished RPC - Authenticate")
	return &pb.AuthenticateResponse{
		User: &pb.User{
			Id:       uint64(res.User.ID),
			Username: res.User.Username,
		},
		Token: res.Token,
	}, nil
}

func (r *rpc) VerifyJWT(ctx context.Context, req *pb.VerifyJWTRequest) (*pb.VerifyJWTResponse, error) {
	res, err := r.userUsecase.VerifyJWT(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	return &pb.VerifyJWTResponse{
		User: &pb.User{
			Id:       uint64(res.ID),
			Username: res.Username,
		},
	}, nil
}

func (r *rpc) UserBorrowBook(ctx context.Context, req *pb.UserBorrowBookRequest) (*pb.UserBorrowBookResponse, error) {
	requestID := utils.ExtractRequestID(ctx)
	userID := utils.ExtractUserID(ctx)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "UserBorrowBook",
		"userID":      userID,
	}).Info("Invoke RPC - UserBorrowBook")
	ctx = context.WithValue(ctx, constants.RequestIDKeyCtx, requestID)

	if req.BookId < 1 || req.UserId < 1 {
		logger.Error("Error while calling method borrowingRecordUsecase.BorrowBook, id must be > 0")
		return nil, errors.New("id must be > 0")
	}

	err := r.borrowingRecordUsecase.BorrowBook(ctx, uint(req.UserId), uint(req.BookId))
	if err != nil {
		logger.WithError(err).Error("Error while calling method borrowingRecordUsecase.BorrowBook")
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "UserBorrowBook",
		"userID":      userID,
	}).Info("Finished RPC - UserBorrowBook")
	return dto.ToUserBorrowBookResponse("successfully borrow a book"), nil
}

func (r *rpc) UserReturnBook(ctx context.Context, req *pb.UserReturnBookRequest) (*pb.UserReturnBookResponse, error) {
	requestID := utils.ExtractRequestID(ctx)
	userID := utils.ExtractUserID(ctx)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "UserReturnBook",
		"userID":      userID,
	}).Info("Invoke RPC - UserReturnBook")
	ctx = context.WithValue(ctx, constants.RequestIDKeyCtx, requestID)

	if req.Id < 1 {
		logger.Error("Error while calling method borrowingRecordUsecase.ReturnBook, id must be > 0")
		return nil, errors.New("id must be > 0")
	}

	err := r.borrowingRecordUsecase.ReturnBook(ctx, uint(req.Id))
	if err != nil {
		logger.WithError(err).Error("Error while calling method borrowingRecordUsecase.ReturnBook")
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "UserReturnBook",
		"userID":      userID,
	}).Info("Finished RPC - UserReturnBook")
	return dto.ToUserReturnBookResponse("successfully return a book"), nil
}
