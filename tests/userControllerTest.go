// user_controller_test.go
package controllers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/d-cryptic/crm-golang-backend/controllers"
	"github.com/gin-gonic/gin"
)

func TestCreateUser(t *testing.T) {
	// Setup Gin router
	router := gin.Default()
	router.POST("/register", controllers.CreateUser)

	// Define test data
	testData := `{"name": "Test User", "email": "test@example.com", "password": "password123"}`
	req, _ := http.NewRequest("POST", "/register", bytes.NewBufferString(testData))
	req.Header.Set("Content-Type", "application/json")

	// Perform request
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check response status code
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.Code)
	}
}