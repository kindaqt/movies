package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestHealthzHandler(t *testing.T) {

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	HealthzHandler(c)

	// Assert Status Code
	assert.Equal(t, http.StatusOK, w.Code)

	var expectedResponse map[string]string = map[string]string{"status": "healthy"}
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedResponse, response)
}
