// rate_limit_middleware_test.go
package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/d-cryptic/crm-golang-backend/middleware"
	"github.com/gin-gonic/gin"
)

func TestRateLimiter(t *testing.T) {
	// Create a new Gin router
	router := gin.New()

	// Apply rate limiting middleware to all routes
	router.Use(middleware.RateLimiter())

	// Define a test endpoint
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	// Create a test request
	req, _ := http.NewRequest("GET", "/test", nil)

	// Create a test response recorder
	w := httptest.NewRecorder()

	// Perform multiple requests within a short period to trigger rate limiting
	for i := 0; i < 20; i++ {
		router.ServeHTTP(w, req)
	}

	// Check if the status code is 429 (Too Many Requests) after exceeding the rate limit
	if w.Code != http.StatusTooManyRequests {
		t.Errorf("Expected status code 429, but got %d", w.Code)
	}
}
