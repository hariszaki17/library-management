package dto

// AuthResponse represents the response structure for getting a user
// @Description User response data
// @Example {"token": "ini token"}
type AuthResponse struct {
	Token string `json:"token"`
}

// AuthResponse represents the response structure for authentication
// @Description Authentication response data
// @Example {"username": "johndoe"}
type AuthRequest struct {
	Username string `json:"username" validate:"required" example:"dudung"`
	Password string `json:"password" validate:"required" example:"maman"`
}

func ToAuthResponse(token string) AuthResponse {
	return AuthResponse{
		Token: token,
	}
}
