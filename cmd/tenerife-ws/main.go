package main

import (
	"net"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/tchaudhry91/tenerife-ws/internal"

	"github.com/gorilla/mux"
)

func main() {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)

	logger.Info("Starting the application")

	port := os.Getenv("PORT")
	if port == "" {
		logger.Errorf("Empty Port Specified")
	}
	router := mux.NewRouter()

	router.HandleFunc("/", internal.WrapLogger(logger, internal.HomeHandler))
	router.HandleFunc("/healthz", internal.WrapLogger(logger, internal.HealthHandler))
	router.HandleFunc("/readyz", internal.WrapLogger(logger, internal.ReadyHandler))

	server := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: router,
	}
	server.ListenAndServe()
}
