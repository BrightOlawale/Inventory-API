package models

// ErrorResponse is a generic response model for all API responses in the application
// This model is used to send validation errors
type ErrorResponse struct {
	ErrorMessage string `json:"errorMessage"`
	Field        string `json:"field"`
}
