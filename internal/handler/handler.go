package handler

import (
	"back-end/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Srv service.SrvMeths
}

func GetHandler(srv service.SrvMeths) *Handler {
	return &Handler{Srv: srv}
}

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
