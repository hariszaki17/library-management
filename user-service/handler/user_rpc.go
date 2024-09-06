// handler/grpc_handler.go
package handler

import (
	"context"

	"github.com/hariszaki17/library-management/proto/constants"
	pb "github.com/hariszaki17/library-management/proto/gen/user/proto"
	"github.com/hariszaki17/library-management/proto/logging"
	"github.com/hariszaki17/library-management/proto/utils"
	"github.com/hariszaki17/library-management/user-service/config"
	"github.com/hariszaki17/library-management/user-service/usecase"
	"github.com/sirupsen/logrus"
)

func NewRPC(useCase usecase.UserUsecase) pb.UserServiceServer {
	return &rpc{
		userUsecase: useCase,
	}
}

type rpc struct {
	userUsecase                       usecase.UserUsecase
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
