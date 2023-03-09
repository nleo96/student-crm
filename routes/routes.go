package routes

import (
	"open-crm-api/controllers"
	"open-crm-api/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func SetupRouter(provider *dig.Container) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(middlewares.HandleErrors)

	AttachUserRoutes(provider, router)

	return router
}

func GetRouterInvoker(destination **gin.Engine) func(*gin.Engine) {
	return func(router *gin.Engine) {
		*destination = router
	}
}

func InvokeRouter(provider *dig.Container) *gin.Engine {
	var router *gin.Engine
	provider.Invoke(GetRouterInvoker(&router))

	return router
}

func AttachUserRoutes(provider *dig.Container, router *gin.Engine) {
	var userController *controllers.UserController
	provider.Invoke(controllers.GetUserControllerInvoker(&userController))

	routerGroup := router.Group("/user")
	{
		routerGroup.POST("/register", userController.SignUp)
		routerGroup.POST("/login", userController.Authenticate)
	}

}
