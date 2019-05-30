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
	diagPort := os.Getenv("DIAG_PORT")
	if diagPort == "" {
		logger.Errorf("Empty Diagnostics Port Specified")
	}
	router := mux.NewRouter()

	routerInternal := mux.NewRouter()

	router.HandleFunc("/", internal.WrapLogger(logger, internal.HomeHandler))
	routerInternal.HandleFunc("/healthz", internal.WrapLogger(logger, internal.HealthHandler))
	routerInternal.HandleFunc("/readyz", internal.WrapLogger(logger, internal.ReadyHandler))

	server := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: router,
	}

	serverInternal := http.Server{
		Addr:    net.JoinHostPort("", diagPort),
		Handler: routerInternal,
	}
	// Do Something Dirty for now
	go serverInternal.ListenAndServe()
	server.ListenAndServe()
}
