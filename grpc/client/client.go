package main

import (
	"context"
	greeter "github.com/masakizk/go/grpc/greeter/protoc"
	"google.golang.org/grpc"
	"log"
)

var client greeter.GreetServiceClient

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	client = greeter.NewGreetServiceClient(conn)
	ctx := context.Background()

	Greet(ctx)
	GreetManyTiles(ctx)
	LongGreet(ctx)
	GreetEveryOne(ctx)
}
