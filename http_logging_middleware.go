package medutils

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Requests with errors are logged using logrus.Error().
// Requests without errors are logged using logrus.Info().
func HttpLogging() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		c.Next()

		end := time.Now().UTC()
		latency := end.Sub(start)

		logFields := make(logrus.Fields)

		logFields["status"] = c.Writer.Status()
		logFields["method"] = c.Request.Method
		logFields["path"] = path
		logFields["ip"] = c.ClientIP()
		logFields["latency"] = latency
		logFields["user-agent"] = c.Request.UserAgent()
		logFields["request-id"] = c.GetString("RequestId")
		logFields["handler"] = c.HandlerName()
		if len(c.Errors) > 0 {
			logFields["errors"] = c.Errors
		}

		entry := logrus.WithFields(logFields)

		if c.Writer.Status() >= 400 {
			logFields["error"] = c.Errors
			entry.Error()
		} else {
			entry.Info()
		}
	}
}
