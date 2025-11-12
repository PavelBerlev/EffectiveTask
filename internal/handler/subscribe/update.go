package subscribe

import (
	"EffectiveTask/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Обновить подписку
// @Description Обновляет поля подписки по ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param request body dto.SubscribeUpdateRequest true "Update request"
// @Router /subscribe [patch]
func (h *Handler) Update(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req dto.SubscribeUpdateRequest
	)

	h.logger.Printf("Handler Update: start remote=%s", c.Request.RemoteAddr)

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Printf("Handler Update: bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	id, statusCode, err := h.subService.Update(ctx, &req)
	if err != nil {
		h.logger.Printf("Handler Update: service error: status=%d id=%d err=%v", statusCode, req.ID, err)
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	h.logger.Printf("Handler Update: success id=%d", id)
	c.JSON(statusCode, gin.H{
		"id": id,
	})
}
