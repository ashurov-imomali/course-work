package route

import (
	"github.com/gin-gonic/gin"
)

func GetRoute(h Routes) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", h.Ping)
	return r
}
