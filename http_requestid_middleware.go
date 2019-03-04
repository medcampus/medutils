package medutils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Create request id with UUID4
		requestID := uuid.New()

		// Expose it for use in the application log
		c.Set("RequestId", requestID.String())

		// Set X-Request-Id header
		c.Header("X-Request-Id", requestID.String())
		c.Next()
	}
}
