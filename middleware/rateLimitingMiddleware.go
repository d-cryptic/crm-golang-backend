// rate_limit_middleware.go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func RateLimiter() gin.HandlerFunc {
	// Create a token bucket rate limiter with a rate of 10 requests per second
	limiter := ratelimit.NewBucketWithRate(10, 10)

	// Return the middleware handler
	return func(c *gin.Context) {
		// Check if the request exceeds the rate limit
		if limiter.TakeAvailable(1) == 0 {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}
		c.Next()
	}
}