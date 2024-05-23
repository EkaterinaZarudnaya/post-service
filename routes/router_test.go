package routes_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ekaterinazarudnaya/post-service/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {

	gin.SetMode(gin.TestMode)
	ginEngine := gin.Default()

	routes.Router(ginEngine)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/welcome", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()

	ginEngine.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "Welcome!"}`, w.Body.String())
}
