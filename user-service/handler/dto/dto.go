package dto

import pb "github.com/hariszaki17/library-management/proto/gen/user/proto"

func ToUserBorrowBookResponse(message string) *pb.UserBorrowBookResponse {
	return &pb.UserBorrowBookResponse{
		Message: message,
	}
}

func ToUserReturnBookResponse(message string) *pb.UserReturnBookResponse {
	return &pb.UserReturnBookResponse{
		Message: message,
	}
}
