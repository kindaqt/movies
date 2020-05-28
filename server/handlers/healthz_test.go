package handlers

import (
	"hello/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func doRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(path, method, nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	return rr
}

func TestHealthzHandler(t *testing.T) {

	router := server.Router().Run()
	w := doRequest(router, "GET", "/healthz")

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned unexpected body: got %v want %v", status, http.StatusOK)
	}

}
