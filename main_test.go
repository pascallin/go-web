package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestServer(t *testing.T) {
	body := gin.H{
		"Status": "Pong",
	}

	router := SetupRouter()

	w := performRequest(router, "GET", "/v1/health-check")

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["Status"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["Status"], value)
}
