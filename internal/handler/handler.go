package handler

import (
	"back-end/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
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

func (h *Handler) GetServices(c *gin.Context) {
	strTypeId := c.Param(":typeId")
	typeId, err := strconv.ParseInt(strTypeId, 10, 64)
	if err != nil {
		h.Log.Error("can't parse 2 int64", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"message":"can't parse 2 int64"})
		return
	}
	services, err := h.Srv.GetServices(typeId)
	if err != nil {
		h.Log.Error("can't get from services", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "wrong mode"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": services})
}
