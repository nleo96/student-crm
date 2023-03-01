package routes

import (
	"context"
	"open-crm-api/controllers"
	"open-crm-api/repositories"
	"open-crm-api/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(context context.Context) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	AttachUserRoutes(context, router)

	return router
}


func AttachUserRoutes(context context.Context, router *gin.Engine) {
	var database *gorm.DB = context.Value("database").(*gorm.DB)

	userRepository := repositories.NewUserRepository(database)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	routerGroup := router.Group("/user")
	{
		routerGroup.POST("/login", userController.Authenticate)
	}

}
