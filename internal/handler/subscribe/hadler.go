package subscribe

import (
	"EffectiveTask/internal/service/subscribe"
	"log"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	api        *gin.Engine
	subService subscribe.SubService
	logger     *log.Logger
}

func NewHandler(api *gin.Engine, subService subscribe.SubService, logger *log.Logger) *Handler {
	return &Handler{
		api:        api,
		subService: subService,
		logger:     logger,
	}
}

func (h *Handler) RouteList() {
	subRoute := h.api.Group("/subscribe")
	subRoute.POST("/", h.Create)
	subRoute.PATCH("/", h.Update)
	subRoute.GET("/:id", h.Get)
	subRoute.DELETE("/:id", h.Delete)
	subRoute.GET("/", h.List)
}
