package server

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/assert/v2"
// )

// func doRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
// 	req, _ := http.NewRequest(method, path, nil)
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)
// 	return w
// }

// func TestHealthzHandler(t *testing.T) {

// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	HealthzHandler(c)
// 	assert.Equal(t, http.StatusOK, w.Code)

// 	// req, err := http.NewRequest("GET", "/healthz", nil)
// 	// assert.Equal(t, err, nil)
// 	// rr := httptest.NewRecorder()

// 	// HealthzHandler.Run(rr, req)

// 	// r.GET("/healthz", HealthzHandler)

// 	// r.ServeHTTP(rr, req)

// 	// assert.Equal(t, rr.Code, http.StatusOK)
// }
