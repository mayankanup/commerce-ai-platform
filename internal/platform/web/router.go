package web

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	app ApplicationInfo,
	logger *slog.Logger,
) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(
		RecoveryMiddleware(logger),
		LoggingMiddleware(logger),
	)

	handler := NewHandler(app)

	RegisterRoutes(router, handler)

	return router
}
