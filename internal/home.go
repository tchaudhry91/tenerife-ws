package internal

import "net/http"

// HomeHandler is the handler for the base path
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Tenerife says hello"))
}
