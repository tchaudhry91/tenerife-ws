package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	// Buffered Interrupt Channel
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Buffered Shutdown Channel
	shutdown := make(chan error, 1)
	server := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: router,
	}
	go func() {
		err := server.ListenAndServe()
		shutdown <- err
	}()
	serverInternal := http.Server{
		Addr:    net.JoinHostPort("", diagPort),
		Handler: routerInternal,
	}

	go func() {
		err := serverInternal.ListenAndServe()
		shutdown <- err
	}()

	select {
	case signalKill := <-interrupt:
		logger.Errorf("Stopping Server: %s", signalKill.String())
	case err := <-shutdown:
		logger.Error(err)
	}

	err := server.Shutdown(context.Background())
	if err != nil {
		logger.Error(err)
	}
	err = serverInternal.Shutdown(context.Background())
	if err != nil {
		logger.Error(err)
	}
}
