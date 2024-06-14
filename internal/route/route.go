package route

import (
	"back-end/internal/handler"
	"github.com/gin-gonic/gin"
)

func GetRoute(h *handler.Handler) *gin.Engine {
	r := gin.Default()
	r.Use(h.CORSMiddleware)
	r.GET("/ping", h.Ping)
	services := r.Group("/services")
	{
		services.GET("/:typeId", h.GetServices)
		services.GET("/by-id/:srvId", h.GetServiceById)
		services.POST("/")
	}
	account := r.Group("/account")
	{
		account.POST("/", h.CreateAccount)
		account.PUT("/", h.AddSum)
	}
	auth := r.Group("/auth")
	{
		auth.POST("/registration", h.Registration)
		auth.POST("/verify", h.Verify)
		auth.POST("/login", h.Login)
	}
	client := r.Group("/client")
	{
		client.GET("/", h.GetClientInfo)
		client.POST("/ident", h.ClientIdent)
		//client.

	}
	return r
}
