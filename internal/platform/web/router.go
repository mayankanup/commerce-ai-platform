package web

import "github.com/gin-gonic/gin"

func NewRouter(app ApplicationInfo) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	handler := NewHandler(app)

	RegisterRoutes(router, handler)

	return router
}
