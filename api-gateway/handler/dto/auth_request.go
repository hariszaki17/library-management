package dto

import (
	pbUser "github.com/hariszaki17/library-management/proto/gen/user/proto"
)

// AuthResponse represents the response structure for getting a user
// @Description User response data
// @Example {"username": "johndoe"}
type AuthResponse struct {
	Username string `json:"username"`
}

// AuthResponse represents the response structure for authentication
// @Description Authentication response data
// @Example {"username": "johndoe"}
type AuthRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func ToAuthResponse(user *pbUser.User) AuthResponse {
	return AuthResponse{
		Username: user.Username,
	}
}
