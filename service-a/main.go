package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	pb "service-a/protos/serviceb/servicebpb"

	"google.golang.org/grpc"
)

const (
	serviceBAddress = "service-b:50051" // service-b address
)

func main() {
	conn, err := grpc.Dial(serviceBAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect to service-b: %v", err)
	}
	defer conn.Close()

	client := pb.NewServiceBClient(conn)

	mux := http.NewServeMux()
	mux.HandleFunc("/call-service-b", func(w http.ResponseWriter, r *http.Request) {
		resp, err := client.CallServiceB(r.Context(), &pb.CallServiceBRequest{})
		if err != nil {
			http.Error(w, "failed to call service-b", http.StatusInternalServerError)
			log.Fatalf("failed to call service-b: %v", err)
		}

		fmt.Fprintf(w, "Service B response: %s", resp.Message)
	})
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		log.Println("starting service a...")
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}()

	stop := make(chan struct{})
	<-stop
	// Block until we receive a signal
	log.Println("service-a is shutting down...")

	// Create a context with a timeout for the shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP server Shutdown: %v", err)
	}

	log.Println("HTTP server shut down gracefully")
}
