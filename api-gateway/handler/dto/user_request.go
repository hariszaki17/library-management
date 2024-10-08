package dto

import (
	pbUser "github.com/hariszaki17/library-management/proto/gen/user/proto"
)

// GetUserResponse represents the response structure for getting a user
// @Description User response data
// @Example {"username": "johndoe"}
type GetUserResponse struct {
	Username string `json:"username"`
}

func ToGetUserResponse(user *pbUser.User) GetUserResponse {
	return GetUserResponse{
		Username: user.Username,
	}
}
