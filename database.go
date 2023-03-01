package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	var dialector gorm.Dialector = postgres.Open(
		fmt.Sprintf(
			"port=%s host=%s dbname=%s user=%s password=%s sslmode=disable",
			os.Getenv("DB_PORT"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
		),
	)
	database, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to setup database connection: %v", err)
	}

	return database
}
