package main

import (
	"context"
	"github.com/tbuchaillot/dkvs/node/server/operations"
	"google.golang.org/grpc"
	"log"
)

func main(){
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := operations.NewChatServiceClient(conn)

	response, err := c.SayHello(context.Background(), &operations.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
}