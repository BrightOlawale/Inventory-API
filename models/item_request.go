package models

import "github.com/go-playground/validator/v10"

// ItemRequest is the request model for creating a new item in the database
type ItemRequest struct {
	Name     string `json:"name" validate:"required"`
	Price    int    `json:"price" validate:"required,gt=0"`
	Quantity int    `json:"quantity" validate:"gte=0"`
}

func (itemInput ItemRequest) ValidateStruct() []*ErrorResponse {
	// Create variable to store validation errors
	var valError []*ErrorResponse

	// Create a validator instance
	validate := validator.New()

	// Validate the itemInput
	err := validate.Struct(itemInput)

	// If error occurred while validating, We will put it in the "valError" variable
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			// Variable to hold current array value
			var element ErrorResponse
			element.ErrorMessage = getErrorMessage(err)
			element.Field = err.Field()
			valError = append(valError, &element)
		}
	}
	// Return validation error
	return valError
}
