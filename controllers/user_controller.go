package controllers

import (
	"net/http"
	"open-crm-api/errors"
	"open-crm-api/models"
	"open-crm-api/services"
	"open-crm-api/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type UserController struct {
	service *services.UserService
}

func (u UserController) Authenticate(context *gin.Context) {
	var credentials models.UserCredentials
	err := context.ShouldBindJSON(&credentials)

	if err != nil {
		context.Error(errors.ErrDataBind.Wrap(err))
		return
	}

	var user *models.User
	user, err = u.service.Authenticate(&credentials)

	if err != nil {
		context.Error(err)
		return
	}

	var token string
	token, err = utils.GenerateJWT(int64(user.ID))

	if err != nil {
		context.Error(err)
		return
	}

	var sessionCookie *http.Cookie = utils.GenerateSessionCookie(token)
	http.SetCookie(context.Writer, sessionCookie)

	context.JSON(http.StatusOK, JSON{
		"message": "Login successful",
		"token":   token,
	})
}

func (u UserController) SignUp(context *gin.Context) {
	var user *models.User = &models.User{}
	err := context.ShouldBindJSON(user)

	if err != nil {
		context.Error(errors.ErrDataBind.Wrap(err))
		return
	}

	user, err = u.service.SignUp(user)

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(
		http.StatusOK,
		JSON{
			"id":        user.ID,
			"name":      user.Name,
			"birthdate": user.Birthdate,
			"email":     user.Email,
		},
	)
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

func GetUserControllerInvoker(
	destination **UserController,
) func(*UserController) {
	return func(controller *UserController) {
		*destination = controller
	}
}

func InvokeUSerController(provider *dig.Container) *UserController {
	var controller *UserController
	provider.Invoke(GetUserControllerInvoker(&controller))

	return controller
}
