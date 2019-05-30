package main

import (
	"net"
	"net/http"
	"os"

	"github.com/tchaudhry91/tenerife-ws/internal"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/", internal.HomeHandler)
	server := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: router,
	}
	server.ListenAndServe()
}
