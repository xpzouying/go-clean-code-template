package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// zap logger middleware
func GinZapLogger() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		if raw != "" {
			path = path + "?" + raw
		}

		zap.S().Named("middleware").Infof("status=%d, latency=%s, client_ip=%s, method=%s, path=%s", statusCode, latency, clientIP, method, path)
	}
}
