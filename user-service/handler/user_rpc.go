// handler/grpc_handler.go
package handler

import (
	"context"

	pb "github.com/hariszaki17/library-management/proto/gen/user/proto"
	"github.com/hariszaki17/library-management/user-service/usecase"
	// Replace with your proto package
)

type GrpcHandler struct {
	userUsecase usecase.UserUsecase
	pb.UnimplementedUserServiceServer
}

func NewGrpcHandler(userUsecase usecase.UserUsecase) *GrpcHandler {
	return &GrpcHandler{userUsecase: userUsecase}
}

func (h *GrpcHandler) GetUser(ctx context.Context, req *pb.GetUserDetailsRequest) (*pb.GetUserDetailsResponse, error) {
	user, err := h.userUsecase.GetUserDetails(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.GetUserDetailsResponse{
		User: &pb.User{
			Username: user.Username,
		},
	}, nil
}
