package dto

// ErrorResponse represents an error response
// @Description Error response structure
// @Example {"message": "Invalid user ID"}
type ErrorResponse struct {
	Message string `json:"message"`
}
