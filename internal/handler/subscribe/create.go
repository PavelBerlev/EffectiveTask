package subscribe

import (
	"EffectiveTask/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Создать подписку
// @Description Создаёт новую подписку (soft-deleted записи игнорируются при проверке существования)
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param request body dto.SubscribeCreateRequest true "Create request"
// @Router /subscribe [post]
func (h *Handler) Create(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req dto.SubscribeCreateRequest
	)

	h.logger.Printf("Handler Create: start remote=%s", c.Request.RemoteAddr)

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Printf("Handler Create: bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	id, statusCode, err := h.subService.Create(ctx, &req)

	if err != nil {
		h.logger.Printf("Handler Create: service error: status=%d user=%s service=%s err=%v", statusCode, req.UserID, req.ServiceName, err)
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	h.logger.Printf("Handler Create: success id=%d user=%s service=%s", id, req.UserID, req.ServiceName)
	c.JSON(statusCode, gin.H{
		"id": id,
	})
}
