package utils

import (
	"context"

	"github.com/hariszaki17/library-management/proto/constants"
	"google.golang.org/grpc/metadata"
)

func ExtractRequestID(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if requestID, ok := md[constants.RequestIDKeyCtx]; ok && len(requestID) > 0 {
			return requestID[0]
		}
	}
	return ""
}

func ExtractUserID(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if userID, ok := md[constants.UserIDKeyCtx]; ok && len(userID) > 0 {
			return userID[0]
		}
	}
	return ""
}
