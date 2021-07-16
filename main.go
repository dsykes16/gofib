package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/dsykes16/gofib/fibonacci"
	pb "github.com/dsykes16/gofib/protos"
	"github.com/dsykes16/gofib/server"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to open port: %s", err)
	}

	fibServer := server.New(fibonacci.LocalMemoizedFibbonacci())

	grpcServer := grpc.NewServer()
	pb.RegisterFibonacciServer(grpcServer, fibServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start grpc server: %s", err)
	}
}
