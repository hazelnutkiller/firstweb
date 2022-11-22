package utils

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// timeout middleware wraps the request context with a timeout
func timeoutMiddleware(timeout time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {

		// wrap the request context with a timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

		defer func() {
			// check if context timeout was reached
			if ctx.Err() == context.DeadlineExceeded {

				// write response and abort the request
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				c.Abort()
			}

			//cancel to clear resources after finished
			cancel()
		}()

		// replace request with context wrapped request
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func Timeout() gin.HandlerFunc {
	// create new gin without any middleware
	engine := gin.New()

	// add timeout middleware with 2 second duration
	engine.Use(timeoutMiddleware(time.Millisecond * 6000))
	    
    if err != nil {
        return err
}
