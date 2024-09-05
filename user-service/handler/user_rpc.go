// handler/grpc_handler.go
package handler

import (
	"context"

	pb "github.com/hariszaki17/library-management/proto/gen/user/proto"
	"github.com/hariszaki17/library-management/user-service/usecase"
)

func NewRPC(useCase usecase.UserUsecase) pb.UserServiceServer {
	return &rpc{
		userUsecase: useCase,
	}
}

type rpc struct {
	userUsecase usecase.UserUsecase
	pb.UnimplementedUserServiceServer // Embed the unimplemented server
}

func (r *rpc) GetUserDetails(ctx context.Context, req *pb.GetUserDetailsRequest) (*pb.GetUserDetailsResponse, error) {
	user, err := r.userUsecase.GetUserDetails(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.GetUserDetailsResponse{
		User: &pb.User{
			Username: user.Username,
		},
	}, nil
}
