package main

import (
	"../proto"
	"./service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

// Start the gRPC server
func main() {

	// Create a new server listener
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	// Create a new proto server
	gRPCServer := grpc.NewServer()

	// Register a new Service Server
	proto.RegisterCalculationServiceServer(gRPCServer, service.InitializeCalculationServer())
	proto.RegisterHashingServiceServer(gRPCServer, service.InitializeHashingServer())

	// Serialize & Deserialize data
	reflection.Register(gRPCServer)

	// Start server listening
	if e := gRPCServer.Serve(listener); e != nil {
		panic(e)
		return
	}
}
