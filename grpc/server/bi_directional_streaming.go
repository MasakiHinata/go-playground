package main

import (
	greeter "github.com/masakizk/go/grpc/greeter/protoc"
	"io"
)

func (*greetServer) GreetEveryone(stream greeter.GreetService_GreetEveryoneServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		_ = stream.Send(&greeter.GreetResponse{
			Result: "Hello " + req.GetGreeting().GetFirstName(),
		})
	}
}
