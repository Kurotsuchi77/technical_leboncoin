package main

import (
	"github.com/Kurotsuchi77/technical_leboncoin/database"
	httpEndpoint "github.com/Kurotsuchi77/technical_leboncoin/endpoint/http"
	"github.com/Kurotsuchi77/technical_leboncoin/fizzbuzz"
	"github.com/sirupsen/logrus"
	"net/http"
)

// MyApp - Struct representing the main application
type MyApp struct {
	endpointHandler *httpEndpoint.Handler
}

// Run - Runs the main application
func (app *MyApp) Run() error {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.Info("Starting Rest API")

	db, err := database.NewDatabase()
	if err != nil {
		logrus.Error("Failed to create database")
		return err
	}
	defer db.Close()

	if res := db.AutoMigrate(&fizzbuzz.Request{}); res.Error != nil {
		logrus.Error("Failed to migrate database")
		return res.Error
	}

	logrus.Info("Database ready")

	app.endpointHandler = httpEndpoint.NewHandler(fizzbuzz.NewService(db))
	app.endpointHandler.SetupRoutes()

	logrus.Info("Endpoint ready")

	if err = http.ListenAndServe(":8080", app.endpointHandler.Router); err != nil {
		logrus.Error("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	app := MyApp{}
	if err := app.Run(); err != nil {
		logrus.Error("Error when starting the application, exiting")
		logrus.Error(err.Error())
	}
}
