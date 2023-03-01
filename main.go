package main

import (
	"context"
	"fmt"
	"open-crm-api/routes"
	"os"
)

func main() {
	database := SetupDatabase()

	apiContext := context.Background()
	apiContext = context.WithValue(apiContext, "database", database)

	router := routes.SetupRouter(apiContext)
	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
