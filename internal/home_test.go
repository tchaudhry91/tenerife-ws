package internal_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tchaudhry91/tenerife-ws/internal"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080", nil)
	w := httptest.NewRecorder()
	internal.HomeHandler(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Incorrect Status Code")
	}
}
