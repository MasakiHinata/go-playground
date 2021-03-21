package main

import (
	"fmt"
	greeter "github.com/masakizk/go/grpc/greeter/protoc"
	"io"
)

func (*greetServer) LongGreet(stream greeter.GreetService_LongGreetServer) error {
	result := "Hello "

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			// we have finished reading the client stream
			return stream.SendAndClose(&greeter.GreetResponse{
				Result: result,
			})
		}
		result += fmt.Sprintf("%s, ", req.Greeting.FirstName)
	}
}
