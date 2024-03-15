package route

import "github.com/gin-gonic/gin"

type Routes interface {
	Ping(c *gin.Context)
}
