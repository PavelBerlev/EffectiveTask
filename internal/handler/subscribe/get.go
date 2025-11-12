package subscribe

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Получить подписку
// @Description Возвращает подписку по ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path int true "Subscribe ID"
// @Router /subscribe/{id} [get]
func (h *Handler) Get(c *gin.Context) {
	ctx := c.Request.Context()

	h.logger.Printf("Handler Get: start remote=%s", c.Request.RemoteAddr)

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.logger.Printf("Handler Get: invalid id: %s", idStr)
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	sub, status, err := h.subService.Get(ctx, id)
	if err != nil {
		h.logger.Printf("Handler Get: service error: status=%d id=%d err=%v", status, id, err)
		c.JSON(status, gin.H{"message": err.Error()})
		return
	}

	h.logger.Printf("Handler Get: success id=%d user=%s service=%s", id, sub.UserID, sub.ServiceName)
	c.JSON(status, sub)
}
