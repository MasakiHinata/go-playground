package main

import (
	greeter "github.com/masakizk/go/grpc/greeter/protoc"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	greeter.RegisterGreetServiceServer(grpcServer, &greetServer{})

	lis, _ := net.Listen("tcp", ":50051")
	log.Println("Start listening at http://localhost:50051")

	err := grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
}

type greetServer struct {
	greeter.UnimplementedGreetServiceServer
}
