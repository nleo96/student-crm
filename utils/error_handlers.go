package utils

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FieldValidationError struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	MinLength int    `json:"minLength,omitempty"`
	MaxLength int    `json:"maxLength,omitempty"`
}

func HandleInvalidRequest(context *gin.Context, err error) {
	var validations validator.ValidationErrors

	if errors.As(err, &validations) {
		failures := make(map[string]*FieldValidationError)
		for _, validation := range validations {
			fieldName := validation.Field()
			fieldValidation, ok := failures[fieldName]

			if !ok {
				fieldValidation = &FieldValidationError{}
				failures[fieldName] = fieldValidation
			}

			switch validation.Tag() {
			case "required":
				fieldValidation.Message = "This field is required"
				fieldValidation.Code = "REQ"

			case "email":
				fieldValidation.Message = "This field must be a valid email address"
				fieldValidation.Code = "INV"

			case "min":
				fieldValidation.Message = "This field must have at least x of length"
				fieldValidation.Code = "MIN"
				fieldValidation.MinLength, _ = strconv.Atoi(validation.Param())

			case "max":
				fieldValidation.Message = "This field must have x of maximum length"
				fieldValidation.Code = "MAX"
				fieldValidation.MaxLength, _ = strconv.Atoi(validation.Param())
			}

		}

		context.AbortWithStatusJSON(
			http.StatusBadRequest,
			JSON{"errors": failures},
		)
		return
	}

	context.AbortWithStatusJSON(
		http.StatusBadRequest,
		JSON{"errors": err.Error()},
	)
}
