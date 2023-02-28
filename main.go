package main

import (
	"fmt"
	"open-crm-api/controllers"
	"open-crm-api/repositories"
	"open-crm-api/services"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var server *gin.Engine

func init() {
	var err error
	var dsn string = fmt.Sprintf(
		"port=%s host=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	server = gin.Default()
	server.Use(cors.Default())
}

func main() {
	userRepository := repositories.NewUserRepository(database)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	server.POST("/login", userController.Authenticate)

	server.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
