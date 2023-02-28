package controllers

import (
	"errors"
	"net/http"
	"open-crm-api/failures"
	"open-crm-api/models"
	"open-crm-api/services"
	"open-crm-api/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func (u UserController) Authenticate(context *gin.Context) {
	var credentials models.UserCredentials
	var err error

	err = context.ShouldBindJSON(&credentials)

	if err != nil {
		context.AbortWithStatusJSON(
			http.StatusBadRequest,
			JSON{"error": err.Error()},
		)
		return
	}

	var user *models.User
	user, err = u.service.Authenticate(&credentials)

	if err != nil {
		if errors.Is(err, failures.ErrInvalidCredentials) {
			context.AbortWithStatusJSON(
				http.StatusUnauthorized,
				JSON{"error": errors.Unwrap(err).Error()},
			)
			return
		}

		context.AbortWithStatusJSON(
			http.StatusInternalServerError,
			JSON{"error": err.Error()},
		)
		return
	}

	var token string
	token, err = utils.GenerateJWT(int64(user.ID))

	if err != nil {
		context.AbortWithStatusJSON(
			http.StatusInternalServerError,
			JSON{"error": err.Error()},
		)
		return
	}

	var sessionCookie *http.Cookie = utils.GenerateSessionCookie(token)
	http.SetCookie(context.Writer, sessionCookie)

	context.JSON(http.StatusOK, JSON{"message": "Login successful"})
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}
