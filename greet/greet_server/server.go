package main

import (
	"context"
	"fmt"
	"github.com/dominik-najberg/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greet function invoked with: %v", req)

	firstName := req.GetGreeting().FirstName
	result := fmt.Sprintf("Hello, %s", firstName)

	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	log.Println("server running...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051") // default port for GRPC
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
