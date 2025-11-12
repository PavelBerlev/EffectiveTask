package subscribe

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Список подписок
// @Description Возвращает список подписок с пагинацией
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param limit query int false "Limit" default(100)
// @Param offset query int false "Offset" default(0)
// @Router /subscribe [get]
func (h *Handler) List(c *gin.Context) {
	ctx := c.Request.Context()

	h.logger.Printf("Handler List: start remote=%s", c.Request.RemoteAddr)

	limitStr := c.DefaultQuery("limit", "100")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 100
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	h.logger.Printf("Handler List: params limit=%d offset=%d", limit, offset)

	subs, status, err := h.subService.List(ctx, limit, offset)
	if err != nil {
		h.logger.Printf("Handler List: service error: status=%d limit=%d offset=%d err=%v", status, limit, offset, err)
		c.JSON(status, gin.H{"message": err.Error()})
		return
	}

	h.logger.Printf("Handler List: success limit=%d offset=%d count=%d", limit, offset, len(subs))
	c.JSON(http.StatusOK, subs)
}
