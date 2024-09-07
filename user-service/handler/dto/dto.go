package dto

import (
	pb "github.com/hariszaki17/library-management/proto/gen/user/proto"
	"github.com/hariszaki17/library-management/user-service/models"
)

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

func ToGetBorrowingCountResponse(borrowingCounts []*models.BorrowingCount) *pb.GetBorrowingCountResponse {
	var res []*pb.BorrowingCount

	for _, bc := range borrowingCounts {
		res = append(res, &pb.BorrowingCount{
			BookId: uint64(bc.BookID),
			Count:  uint64(bc.Count),
		})
	}

	return &pb.GetBorrowingCountResponse{
		BorrowingCount: res,
	}
}
