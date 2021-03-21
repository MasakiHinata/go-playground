package main

import (
	"context"
	"fmt"
	greeter "github.com/masakizk/go/grpc/greeter/protoc"
	"time"
)

var requests = []*greeter.GreetRequest{
	{Greeting: &greeter.Greeting{
		FirstName: "Stephane",
		LastName:  "",
	}},
	{Greeting: &greeter.Greeting{
		FirstName: "John",
		LastName:  "",
	}},
	{Greeting: &greeter.Greeting{
		FirstName: "Lucy",
		LastName:  "",
	}},
}

func LongGreet(ctx context.Context) {
	stream, _ := client.LongGreet(ctx)

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		_ = stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, _ := stream.CloseAndRecv()
	fmt.Printf("LongGreet Response: %v", res.Result)
}
