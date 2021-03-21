package main

import (
	"fmt"
	greeter "github.com/masakizk/go/grpc/greeter/protoc"
	"time"
)

func (g *greetServer) GreetManyTimes(req *greeter.GreetRequest, stream greeter.GreetService_GreetManyTimesServer) error {

	for i := 0; i < 10; i++ {
		res := &greeter.GreetResponse{
			Result: fmt.Sprintf("Hello %s: %d", req.Greeting.GetFirstName(), i),
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}
