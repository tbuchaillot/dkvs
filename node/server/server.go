package server

import (
	"errors"
	"fmt"
	"github.com/tbuchaillot/dkvs/node/databases"
	"github.com/tbuchaillot/dkvs/node/server/operations"
	"log"
	"net"
	"google.golang.org/grpc"
)

type Server struct{
	listener net.Listener
	grpcServer *grpc.Server
	db databases.Database
}

func NewServer(db databases.Database) (*Server,error){
	newServer := &Server{
		db : db,
	}

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
	operationServer := operations.NewOperationService(srv.db)
	operations.RegisterOperationServiceServer(srv.grpcServer, operationServer)
}

func (srv *Server) Serve() error{
	log.Printf("Serving :D ")
	return srv.grpcServer.Serve(srv.listener)
}
