package handler

import "back-end/internal/service"

type Handler struct {
	Srv service.SrvMeths
}

func GetHandler(srv service.SrvMeths) *Handler {
	return &Handler{Srv: srv}
}

func (h *Handler) Ping(c *gin.Context) 
}
