package models

import "github.com/go-playground/validator/v10"

// getErrorMessage: the error message is returned based on the validation rule using the switch case condition selection
func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		// Return the error message if validation fails for the required rule.
		return err.Field() + " is required"
	case "gt":
		// gt means a specific input must be greater than the specified value in the validation rule.
		return "the value of " + err.Field() + " must be greater than " + err.Param()
	case "gte":
		// gte means a specific input must be greater than or equal to the specified value in the validation rule.
		return "the value of " + err.Field() + " must be greater than or equals  " + err.Param()
	default:
		// Return the default error message.
		return "Validation error in " + err.Field()
	}
}
