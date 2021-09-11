package main

import (
	"net/http"

	"github.com/Sunshine31/financial-app/internal/api"
	"github.com/Sunshine31/financial-app/internal/config"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.WithField("version", config.Version).Debug("Starting server.")
	router, err := api.NewRouter()
	if err != nil {
		logrus.New().WithError(err).Fatal("Error building router")
	}
	const addr = "0.0.0.0:8088"
	server := http.Server{
		Handler: router,
		Addr:    addr,
	}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Error("Server failed.")
	}
}
