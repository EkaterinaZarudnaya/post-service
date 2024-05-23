package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ekaterinazarudnaya/post-service/routes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetPosts(t *testing.T) {

	gin.SetMode(gin.TestMode)

	ginEngine := gin.Default()
	routes.Router(ginEngine)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/posts/", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()

	ginEngine.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetPostById(t *testing.T) {

	gin.SetMode(gin.TestMode)

	ginEngine := gin.Default()
	routes.Router(ginEngine)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/posts/2", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()

	ginEngine.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestNewPost(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ginEngine := gin.Default()

	routes.Router(ginEngine)

	postData := map[string]string{
		"user_email": "test_user@gmail.com",
		"title":      "Test user title",
		"content":    "The content of the test user post",
	}

	jsonData, err := json.Marshal(postData)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/posts/", bytes.NewBuffer(jsonData))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	ginEngine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdatePostById(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ginEngine := gin.Default()
	routes.Router(ginEngine)

	postData := map[string]string{
		"user_email": "update_test_user@gmail.com",
		"title":      "Test user title",
		"content":    "The content of the test user post",
	}

	jsonData, err := json.Marshal(postData)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/posts/2", bytes.NewBuffer(jsonData))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	ginEngine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeletePostByIdNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ginEngine := gin.Default()
	routes.Router(ginEngine)

	req, err := http.NewRequest(http.MethodDelete, "/api/v1/posts/999", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()

	ginEngine.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeletePostById(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ginEngine := gin.Default()
	routes.Router(ginEngine)

	// TODO: change ID
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/posts/20", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()

	ginEngine.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestNewPostBindJSONError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ginEngine := gin.Default()
	routes.Router(ginEngine)

	invalidJsonData := `{"user_email": "test_user@gmail.com", "title": "Test user title", "content": "The content of the test user post"`

	req, err := http.NewRequest(http.MethodPost, "/api/v1/posts/", bytes.NewBufferString(invalidJsonData))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	ginEngine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestNewPostValidationError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ginEngine := gin.Default()
	routes.Router(ginEngine)

	postData := map[string]string{
		"user_email": "invalid-email",
		"title":      "",
		"content":    "",
	}

	jsonData, err := json.Marshal(postData)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/posts/", bytes.NewBuffer(jsonData))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	ginEngine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdatePostByIdBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ginEngine := gin.Default()
	routes.Router(ginEngine)

	postData := map[string]string{
		"user_email": "invalid-email",
		"title":      "Test user title",
		"content":    "The content of the test user post",
	}

	jsonData, err := json.Marshal(postData)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/posts/2", bytes.NewBuffer(jsonData))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	ginEngine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateInvalidPostIDFormat(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ginEngine := gin.Default()
	routes.Router(ginEngine)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/posts/1е", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	ginEngine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateInvalidJSONFormat(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ginEngine := gin.Default()
	routes.Router(ginEngine)
	postData := map[string]string{
		"user_email": "invalid-email",
		"title":      "",
		"content":    "",
	}

	jsonData, err := json.Marshal(postData)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/posts/2", bytes.NewBuffer(jsonData))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	ginEngine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateBindJSONError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ginEngine := gin.Default()
	routes.Router(ginEngine)

	invalidJsonData := `{"user_email": "test_user@gmail.com", "title": "Test user title", "content": "The content of the test user post"`
	jsonData, _ := json.Marshal(invalidJsonData)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/posts/2", bytes.NewBuffer(jsonData))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	ginEngine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetInvalidPostIDFormat(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ginEngine := gin.Default()
	routes.Router(ginEngine)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/posts/invalid_id", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	ginEngine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
