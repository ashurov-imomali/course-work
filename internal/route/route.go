package route

import (
	"back-end/internal/handler"
	"github.com/gin-gonic/gin"
)

func GetRoute(h *handler.Handler) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", h.Ping)
	services := r.Group("/services")
	services.GET("/:typeId", h.GetServices)
	return r
}
