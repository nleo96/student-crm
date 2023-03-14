package controllers

import (
	"net/http"
	"open-crm-api/models"
	"open-crm-api/repositories"
	"open-crm-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type ContactController struct {
	service *services.ContactService
}

func (c ContactController) FindById(context *gin.Context) {
	id, _ := strconv.ParseUint(context.Param("id"), 10, 32)

	contact, err := c.service.Repository().FindByID(
		uint(id),
		repositories.ScanModeDeep,
	)

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, contact)
}

func (c ContactController) Create(context *gin.Context) {
	contact := &models.Contact{}

	err := context.ShouldBindJSON(contact)
	if err != nil {
		context.Error(err)
		return
	}

	contact, err = c.service.Repository().Create(contact)
	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, contact)
}

func NewContactController(service *services.ContactService) *ContactController {
	return &ContactController{service: service}
}

func GetContactControllerInvoker(
	destination **ContactController,
) func(*ContactController) {
	return func(controller *ContactController) {
		*destination = controller
	}
}

func InvokeContactController(provider *dig.Container) *ContactController {
	var controller *ContactController
	provider.Invoke(GetContactControllerInvoker(&controller))

	return controller
}
