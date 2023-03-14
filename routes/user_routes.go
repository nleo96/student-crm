package routes

import (
	"open-crm-api/controllers"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func AttachUserRoutes(provider *dig.Container, router *gin.Engine) {
	userController := controllers.InvokeUSerController(provider)

	routerGroup := router.Group("/user")
	{
		routerGroup.POST("/register", userController.SignUp)
		routerGroup.POST("/login", userController.Authenticate)
	}
}
