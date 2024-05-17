package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func StartListen(handler http.Handler, srv *Server) error {
	var err error
	go func(err *error) {
		*err = http.ListenAndServe(fmt.Sprintf("%s:%d", srv.Host, srv.Port), handler)
	}(&err)
	return err
}
