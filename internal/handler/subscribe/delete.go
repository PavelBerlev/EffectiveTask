package subscribe

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Удалить подписку
// @Description Soft-delete подписки по ID (устанавливает deleted_at)
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path int true "Subscribe ID"
// @Success 200 {object} map[string]int64 "id"
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /subscribe/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	h.logger.Printf("Handler Delete: start remote=%s", c.Request.RemoteAddr)

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.logger.Printf("Handler Delete: invalid id: %s", idStr)
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	deletedID, status, err := h.subService.Delete(ctx, id)
	if err != nil {
		h.logger.Printf("Handler Delete: service error: status=%d id=%d err=%v", status, id, err)
		c.JSON(status, gin.H{"message": err.Error()})
		return
	}

	h.logger.Printf("Handler Delete: success id=%d", deletedID)
	c.JSON(status, gin.H{"id": deletedID})
}
