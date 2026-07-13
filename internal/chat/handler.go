package chat

import (
	"net/http"

	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/mayankanup/commerce-ai-platform/internal/agent"
)

type Handler struct {
	agent  *agent.Agent
	logger *slog.Logger
}

func NewHandler(
	agent *agent.Agent,
	logger *slog.Logger,
) *Handler {
	return &Handler{
		agent:  agent,
		logger: logger,
	}
}

func (h *Handler) Handle(c *gin.Context) {

	var request ChatRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	result, err := h.agent.Chat(
		c.Request.Context(),
		request.Message,
	)

	if err != nil {
		h.logger.Error(
			"chat failed",
			"error", err,
		)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, ChatResponse{
		Response: result.Response,
	})
}
