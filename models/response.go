package models

// Response is a generic response model for all API responses in the application
// This model can be used to send validation errors, success messages, etc
// The "T any" syntax is used to make the "Data" field generic
type Response[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}
