package server

import (
	"errors"
	"fmt"
	"github.com/tbuchaillot/dkvs/node/server/operations"
	"log"
	"net"
	"google.golang.org/grpc"
)

type Server struct{
	listener net.Listener
	grpcServer *grpc.Server
}

func NewServer() (*Server,error){
	newServer := &Server{}

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		return nil,errors.New(fmt.Sprintf("failed to listen: %v", err))
	}

	newServer.listener = lis


	grpcServer := grpc.NewServer()
	newServer.grpcServer = grpcServer

	newServer.registerServices()

	return newServer,nil
}

func (srv *Server) registerServices(){
	operationServer := operations.Server{}
	operations.RegisterChatServiceServer(srv.grpcServer, &operationServer)

}

func (srv *Server) Serve() error{
	log.Printf("Serving :D ")
	return srv.grpcServer.Serve(srv.listener)
}
