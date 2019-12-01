package main

import (
	"awesomeProject/finance-app-backend/backend/internal/api"
	"awesomeProject/finance-app-backend/backend/internal/config"
	"awesomeProject/finance-app-backend/backend/internal/database"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.WithField("version", config.Version).Debug("Starting server")

	// Creating new database
	db, err := database.New()
	if err != nil {
		logrus.WithError(err).Fatal()
	}

	// Creating new router
	router, err := api.NewRouter(db)
	if err != nil {
		logrus.WithError(err).Fatal("Error building router")
	}
	const addr = "0.0.0.0:8080"
	server := http.Server{
		Addr:              addr,
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Error("Server failed.")
	}
}
