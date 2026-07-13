package chat

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	router *gin.Engine,
	handler *Handler,
) {
	router.POST(
		"/api/v1/chat",
		handler.Handle,
	)
}
