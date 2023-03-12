package main

import (
	"crypto/tls"
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

	cert, err := tls.LoadX509KeyPair(".ssl/cert.pem", ".ssl/key.pem")
	if err != nil {
		panic(err)
	}

	return &Application{
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")),
			Handler: routes.InvokeRouter(provider),
			TLSConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
			},
		},
	}
}

func (a Application) Run() error {
	return a.server.ListenAndServe()
}
