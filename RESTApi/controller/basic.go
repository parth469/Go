package controller

import (
	"strings"

	validator "github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func formatValidationErrors(err error) map[string]string {
	errors := make(map[string]string)

	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()
		tag := err.Tag()

		var message string
		switch tag {
		case "required":
			message = field + " is required"
		case "email":
			message = field + " must be a valid email"
		case "min":
			message = field + " must be at least " + err.Param() + " characters long"
		case "max":
			message = field + " must be at most " + err.Param() + " characters long"
		default:
			message = field + " is invalid"
		}

		errors[strings.ToLower(field)] = message
	}
	return errors
}

func CreateValidator() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}
