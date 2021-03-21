package main

import (
	"context"
	greeter "github.com/masakizk/go/grpc/greeter/protoc"
	"log"
)

func Greet(ctx context.Context) {
	res, _ := client.Greet(ctx, &greeter.GreetRequest{
		Greeting: &greeter.Greeting{
			FirstName: "Alice",
			LastName:  "Smith",
		},
	})
	log.Println(res.Result)
}
