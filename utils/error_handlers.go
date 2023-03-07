package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleBadRequestError(context *gin.Context, err error) {
	context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
}
