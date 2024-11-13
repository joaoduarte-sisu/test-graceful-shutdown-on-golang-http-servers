package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "service-b/protos/serviceb/servicebpb"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnsafeServiceBServer
}

func (s *server) CallServiceB(ctx context.Context, req *pb.CallServiceBRequest) (*pb.CallServiceBResponse, error) {
	time.Sleep(20 * time.Second)
	return &pb.CallServiceBResponse{
		Message: "hello",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterServiceBServer(grpcServer, &server{})

	log.Printf("service-b listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
