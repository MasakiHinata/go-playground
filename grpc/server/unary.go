package main

import (
	"context"
	"fmt"
	greeter "github.com/masakizk/go/grpc/greeter/protoc"
)

func (g *greetServer) Greet(ctx context.Context, request *greeter.GreetRequest) (*greeter.GreetResponse, error) {
	return &greeter.GreetResponse{
		Result: fmt.Sprintf("Hello %s!", request.Greeting.FirstName),
	}, nil
}
