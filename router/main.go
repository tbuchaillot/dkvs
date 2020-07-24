package main

import (
	"github.com/tbuchaillot/dkvs/router/server"
)

func main(){
	srv := server.NewServer("",8080)
	srv.RegiterRoutes()
	srv.Serve()
}