package main

import (
	"context"
	greeter "github.com/masakizk/go/grpc/greeter/protoc"
	"io"
	"log"
)

func GreetManyTiles(ctx context.Context) {
	req := &greeter.GreetRequest{
		Greeting: &greeter.Greeting{
			FirstName: "Alice",
			LastName:  "Smith",
		},
	}

	stream, _ := client.GreetManyTimes(ctx, req)

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break // we've reached the end of the stream.
		}
		log.Printf("Response from Server Streaming: %v", msg.GetResult())
	}
}
