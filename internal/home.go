package internal

import (
	"encoding/json"
	"net/http"
)

// HomeResponse is the response of the HomeController
type HomeResponse struct {
	Value string `json:"value,omitempty"`
}

// HomeHandler is the handler for the base path
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	home := HomeResponse{
		Value: "All Good!",
	}
	resp, err := json.Marshal(home)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(resp)
}
