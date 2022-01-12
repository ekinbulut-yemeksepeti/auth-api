package main

import (
	"net/http"

	"github.com/ekinbulut-yemeksepeti/auth-api/internal/authentication"
	"github.com/ekinbulut-yemeksepeti/auth-api/internal/token"

	transportHTTP "github.com/ekinbulut-yemeksepeti/auth-api/internal/transport/http"
	log "github.com/sirupsen/logrus"
)

const (
	name    = "Auth api"
	version = "v1.0"
)

type App struct {
	name    string
	version string
}

func (app *App) Run() error {

	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    app.name,
			"AppVersion": app.version,
		}).Info("Setting Up Our APP")

	// initilize services
	tokenService := token.NewService()
	authService := authentication.NewService(tokenService)

	// creating handler with dependencies
	handler := transportHTTP.NewHandler(authService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Error("Failed to set up server")
		return err
	}

	return nil
}

func main() {

	app := App{
		name:    name,
		version: version,
	}

	if err := app.Run(); err != nil {
		log.Error(err)
		log.Fatal("Error starting up our REST API")
	}

}
