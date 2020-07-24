package main

import (
	"flag"
	"github.com/tbuchaillot/dkvs/node/config"
	"github.com/tbuchaillot/dkvs/node/databases"
	"github.com/tbuchaillot/dkvs/node/server"
	"log"
)

var (
	configPath = flag.String("config","default.test.config","Config file path for dkvs node.")
)

func main(){

	if err := config.ParseConfg(*configPath); err != nil {
		log.Fatalf("Unable to load config file %v", *configPath)
	}

	db, err := databases.NewDatabase(databases.BOLT_TYPE,config.Global().Name, config.Global().Extension);
	if err != nil {
		log.Fatalf("Unable to initiliaze db %v.%v , error %v", config.Global().Name,config.Global().Extension, err)
	}
	defer db.Close()

	srv, err := server.NewServer(db)
	if err != nil {
		log.Fatalf("Unable to initiliaze server error %v", err)

	}

	srv.Serve()
}