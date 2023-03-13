package controllers

import (
	"open-crm-api/services"

	"go.uber.org/dig"
)

type ContactController struct {
	service *services.ContactService
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
