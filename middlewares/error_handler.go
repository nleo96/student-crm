package middlewares

import (
	"net/http"
	"open-crm-api/errors"

	"github.com/gin-gonic/gin"
)

func HandleErrors(context *gin.Context) {
	context.Next()

	if len(context.Errors) > 0 {
		wrapedError := context.Errors.Last().Err
		var err *errors.Error

		if errors.As(wrapedError, err) {
			httpError := errorToHttpError(err)
			context.AbortWithStatusJSON(
				httpError.Code(),
				JSON{"message": httpError.Error()},
			)
		} else {
			context.AbortWithStatusJSON(
				http.StatusInternalServerError,
				JSON{"message": "a internal server error occur"},
			)
		}
	}
}

func errorToHttpError(err *errors.Error) *errors.HTTPError {
	switch err.Kind() {
	case errors.ErrKindDataNotFound:
		return errors.NewHTTPError(http.StatusNotFound, err.Error(), err)

	case errors.ErrKindInvalidCredentials:
		return errors.NewHTTPError(http.StatusUnauthorized, err.Error(), err)

	default:
		return errors.NewHTTPError(http.StatusInternalServerError, "unknown internal server error", err)
	}
}
