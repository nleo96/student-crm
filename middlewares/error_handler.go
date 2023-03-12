package middlewares

import (
	"fmt"
	"net/http"
	"open-crm-api/errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
)

func HandleErrors(context *gin.Context) {
	context.Next()

	if len(context.Errors) > 0 {
		wrapedError := context.Errors.Last()
		var err *errors.Error
		var httpError *errors.HTTPError

		if errors.As(wrapedError, &err) {
			httpError = HTTPErrorFrom(err)
		} else {
			httpError = errors.NewUnkownHTTPError(wrapedError)
		}

		context.AbortWithStatusJSON(httpError.Code(), httpError)
	}
}

func HTTPErrorFrom(err *errors.Error) *errors.HTTPError {
	switch err.Kind() {
	case errors.ErrKindDataNotFound:
		return errors.NewHTTPError(http.StatusNotFound, err)

	case errors.ErrKindInvalidCredentials:
		return errors.NewHTTPError(http.StatusUnauthorized, err)

	case errors.ErrKindDataBind:
		return errors.NewHTTPError(
			http.StatusBadRequest,
			err,
		).WithCause(GetCauseOfDataBindError(err))

	default:
		return errors.NewUnkownHTTPError(err)
	}
}

func GetCauseOfDataBindError(
	err *errors.Error,
) map[string]*errors.DataValidationError {
	var validations validator.ValidationErrors

	if errors.As(err, &validations) {
		failures := make(map[string]*errors.DataValidationError)
		for _, validation := range validations {
			fieldName := strcase.ToLowerCamel(validation.Field())
			fieldValidation, ok := failures[fieldName]

			if !ok {
				fieldValidation = &errors.DataValidationError{}
				failures[fieldName] = fieldValidation
			}

			validationValue, _ := strconv.Atoi(validation.Param())

			switch validation.Tag() {
			case "required":
				fieldValidation.Code = "REQ"
				fieldValidation.Message = "This field is required"

			case "email":
				fieldValidation.Code = "INV"
				fieldValidation.Message = "This field must be a valid email address"

			case "min":
				fieldValidation.Code = "MIN"
				fieldValidation.Message = fmt.Sprintf(
					"This field must have at least %v of length",
					validationValue,
				)
				fieldValidation.Value = validationValue

			case "max":
				fieldValidation.Code = "MAX"
				fieldValidation.Message = fmt.Sprintf(
					"This field must have %v of maximum length",
					validationValue,
				)
				fieldValidation.Value = validationValue
			}
		}

		return failures
	}
	return nil
}
