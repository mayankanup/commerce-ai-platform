package web

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		duration := time.Since(start)

		logger.Info(
			"HTTP request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"client_ip", c.ClientIP(),
			"duration_ms", duration.Milliseconds(),
		)
	}
}

func RecoveryMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered any) {

		logger.Error(
			"Panic recovered",
			"error", recovered,
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
		)

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "internal server error",
			},
		)
	})
}
