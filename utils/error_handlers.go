package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleInvalidRequest(context *gin.Context, err error) {
	var validations validator.ValidationErrors

	if errors.As(err, &validations) {
		var failures map[string][]string = make(map[string][]string)
		for _, validation := range validations {
			var message string
			switch validation.Tag() {
			case "required":
				message = "is mandatory"

			case "email":
				message = "must be a valid email address"

			case "min":
				message = fmt.Sprintf(
					"the length or value must be at least %s",
					validation.Param(),
				)

			case "max":
				message = fmt.Sprintf(
					"the length or value must be a maximum of %s",
					validation.Param(),
				)
			}

			fieldName := validation.Field()
			failures[fieldName] = append(failures[fieldName], message)
		}

		context.AbortWithStatusJSON(
			http.StatusBadRequest,
			JSON{"errors": failures},
		)
	}

	context.AbortWithStatusJSON(
		http.StatusBadRequest,
		JSON{"errors": err.Error()},
	)
}
