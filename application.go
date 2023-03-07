package main

import (
	"fmt"
	"net/http"
	"open-crm-api/provider"
	"open-crm-api/routes"
	"os"
)

type Application struct {
	server *http.Server
}

func SetupApplication() *Application {
	provider := provider.SetupProvider()

	return &Application{
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")),
			Handler: routes.InvokeRouter(provider),
		},
	}
}

func (a Application) Run() error {
	return a.server.ListenAndServe()
}
