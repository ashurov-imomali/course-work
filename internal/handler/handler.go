package handler

import (
	"back-end/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	Srv service.SrvMeths
	Log *zap.Logger
}

func GetHandler(srv service.SrvMeths, zLog *zap.Logger) *Handler {
	return &Handler{Srv: srv, Log: zLog}
}

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
