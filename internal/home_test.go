package internal_test

import (
	"io/ioutil"
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
	if data, _ := ioutil.ReadAll(w.Body); string(data) != "Tenerife says hello" {
		t.Errorf("Incorrect data received:%s", string(data))
	}
}
