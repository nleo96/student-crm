package provider

import (
	"open-crm-api/controllers"
	"open-crm-api/database"
	"open-crm-api/repositories"
	"open-crm-api/routes"
	"open-crm-api/services"

	"go.uber.org/dig"
)

func SetupProvider() *dig.Container {
	provider := dig.New()

	provider.Provide(func() *dig.Container { return provider })

	provider.Provide(database.SetupDatabase)
	provider.Provide(routes.SetupRouter)

	provider.Provide(repositories.NewUserRepository)

	provider.Provide(services.NewUserService)

	provider.Provide(controllers.NewUserController)

	return provider
}
