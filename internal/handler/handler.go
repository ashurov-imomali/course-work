package handler

import (
	"back-end/internal/service"
	"back-end/pkg/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
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

func (h *Handler) CORSMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
	}
}

func (h *Handler) GetServices(c *gin.Context) {
	strTypeId := c.Param("typeId")
	typeId, err := strconv.ParseInt(strTypeId, 10, 64)
	if err != nil {
		h.Log.Error("can't parse 2 int64", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"message": "can't parse 2 int64"})
		return
	}
	srvName := c.Query("name")
	strId := c.Query("id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		h.Log.Info("can't parse id", zap.Error(err))
	}
	services, err := h.Srv.GetServices(typeId, id, srvName)
	if err != nil {
		h.Log.Error("can't get from services", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "wrong mode"})
		return
	}
	c.JSON(http.StatusOK, services)
}

func (h *Handler) Login(c *gin.Context) {
	var loginInfo models.Login
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		h.Log.Error("can't parse 2 struct")
		c.JSON(http.StatusBadRequest, gin.H{"message": "can't parse 2 struct"})
		return
	}
	token, sErr := h.Srv.Login(&loginInfo)
	if sErr.Err != nil {
		h.Log.Error(sErr.Message, zap.Error(sErr.Err))
		c.JSON(sErr.Status, gin.H{"message": sErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) Registration(c *gin.Context) {
	var newClient models.Client
	if err := c.ShouldBindJSON(&newClient); err != nil {
		h.Log.Info("can't parse 2 struct. Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"message": "can't parse 2 struct"})
		return
	}
	if err := h.Srv.Registration(&newClient); err.Err != nil {
		h.Log.Error("Registration error: ", zap.Error(err.Err))
		c.JSON(err.Status, gin.H{"message": err.Message})
		return
	}
	log.Println("ok")
	c.JSON(http.StatusCreated, newClient)
}

func (h *Handler) GetServiceById(c *gin.Context) {
	strSrvId := c.Param("srvId")
	srvId, err := strconv.ParseInt(strSrvId, 10, 64)
	if err != nil {
		h.Log.Error("can't parse 2 int64", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"message": "can't parse 2 int64"})
		return
	}
	srv, err := h.Srv.GetService(srvId)
	if err != nil {
		h.Log.Error("error get service", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't get services"})
		return
	}
	c.JSON(http.StatusOK, srv)
}

func (h *Handler) CreateAccount(c *gin.Context) {

}

func (h *Handler) AddSum(c *gin.Context) {

}

func (h *Handler) GetClientInfo(c *gin.Context) {

}

func (h *Handler) ClientIdent(c *gin.Context) {

}

func (h *Handler) Verify(c *gin.Context) {
	var confirm models.VerifyLogin
	if err := c.ShouldBindJSON(&confirm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "couldn't parse 2 json"})
		return
	}
	if hErr := h.Srv.VerifyLogin(&confirm); hErr.Err != nil {
		h.Log.Error("[VERIFY]", zap.Error(hErr.Err))
		c.JSON(hErr.Status, gin.H{"message": hErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
