package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"post-service/routes"
)

func TestPingRoute(t *testing.T) {

	gin.SetMode(gin.TestMode)
	ginEngine := gin.New()
	ginEngine = gin.Default()

	routes.Router(ginEngine)

	req, err := http.NewRequest(http.MethodGet, "/welcome", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()

	ginEngine.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "Welcome!"}`, w.Body.String())
}
