package subscribe

import (
	"EffectiveTask/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Подсчитать суммарную стоимость подписок за период
// @Description Возвращает суммарную цену подписок за период. Фильтрация по user_id ИЛИ service_name (обязательно один из них).
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param request body dto.SubscribeTotalCostRequest true "Total cost request"
// @Router /subscribe/total-cost [post]
func (h *Handler) CalculateTotalCost(c *gin.Context) {
	ctx := c.Request.Context()

	h.logger.Printf("Handler CalculateTotalCost: start remote=%s", c.Request.RemoteAddr)

	var req dto.SubscribeTotalCostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Printf("Handler CalculateTotalCost: bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	totalCost, status, err := h.subService.CalculateTotalCost(ctx, &req)
	if err != nil {
		h.logger.Printf("Handler CalculateTotalCost: service error: status=%d user_id=%v service_name=%v err=%v", status, req.UserID, req.ServiceName, err)
		c.JSON(status, gin.H{"message": err.Error()})
		return
	}

	h.logger.Printf("Handler CalculateTotalCost: success total_cost=%d", totalCost)
	c.JSON(status, gin.H{"total_cost": totalCost})
}
