package handler

import (
	"back-end/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

type Handler struct {
	Srv service.SrvMeths
}

func GetHandler(srv service.SrvMeths) *Handler {
	return &Handler{Srv: srv}
}

func (h *Handler) Ping(c *gin.Context) {
	defer func() {
		go func() {
			time.Sleep(2 * time.Second)
			err := os.RemoveAll("app-log.log")
			if err != nil {
				log.Println(err)
				return
			}
		}()
	}()
	c.JSON(200, gin.H{"message": "pong"})

}
