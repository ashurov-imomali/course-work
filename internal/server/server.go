package server

import (
	"back-end/internal/handler"
)

type Server struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func StartListen(handler handler.Handler, server *Server) {
	srv.ListenAndServe()
}
