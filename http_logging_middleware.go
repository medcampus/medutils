package medutils

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Requests with errors are logged using logrus.Error().
// Requests without errors are logged using logrus.Info().
func Http_Logging() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		c.Next()

		end := time.Now().UTC()
		latency := end.Sub(start)

		entry := logrus.WithFields(logrus.Fields{
			"status":     c.Writer.Status(),
			"method":     c.Request.Method,
			"path":       path,
			"ip":         c.ClientIP(),
			"latency":    latency,
			"user-agent": c.Request.UserAgent(),
			"request-id": c.GetHeader("X-Request-Id"),
			"time":       end.Format(time.RFC3339),
		})

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			entry.Error(c.Errors.String())
		} else {
			entry.Info()
		}
	}
}
