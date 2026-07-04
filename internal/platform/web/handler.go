package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	app ApplicationInfo
}

func NewHandler(app ApplicationInfo) *Handler {
	return &Handler{
		app: app,
	}
}

func (h *Handler) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"application": h.app.Name,
		"version":     h.app.Version,
		"environment": h.app.Environment,
		"status":      "UP",
	})
}

func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}

func (h *Handler) Ready(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "READY",
	})
}
