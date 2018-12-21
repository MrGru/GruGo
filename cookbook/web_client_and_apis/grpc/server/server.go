package main

import (
	"fmt"
	"net"

	"github.com/MrGru/GruGo/cookbook/web_client_and_apis/grpc/greeter"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	greeter.RegisterGreeterServiceServer(grpcServer, &Greeter{Exclaim: true})
	lis, err := net.Listen("tcp", ":4444")
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port :4444")
	grpcServer.Serve(lis)
}
