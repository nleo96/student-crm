package routes

import (
	"open-crm-api/controllers"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func AttachContactRoutes(provider *dig.Container, router *gin.Engine) {
	contactController := controllers.InvokeContactController(provider)

	routerGroup := router.Group("/contact")
	{
		routerGroup.GET("/find/:id", contactController.FindById)
		routerGroup.POST("", contactController.Create)
	}
}
