package server

import (
	"fmt"
	"github.com/tbuchaillot/dkvs/router/clients"
	"github.com/tbuchaillot/dkvs/router/management"

	"github.com/labstack/echo/v4"

)

type Server struct {
	nodes     *management.Nodes
	srv		*echo.Echo
	cl *clients.Client

	host string
	port int
}

func NewServer(host string, port int) *Server {
	newSrv := &Server{host: host,port: port}

	newSrv.nodes = management.NewNodes()
	newSrv.srv = echo.New()
	newSrv.cl =clients.NewClient()

	return newSrv
}

func (s *Server) RegiterRoutes(){
	s.srv.POST("/keys/:key",s.SetKey)
	s.srv.GET("/keys/:key",s.GetKey)

	s.srv.POST("/nodes",s.RegisterNode)
}


func (s *Server) Serve(){
	s.srv.Start(s.host+":"+fmt.Sprint(s.port))
}