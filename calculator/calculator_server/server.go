package main

import (
	"context"
	"fmt"
	"github.com/dominik-najberg/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (s *server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Printf("Sum function invoked with: %v", req)

	firstNumber := req.Sum.FirstNumber
	secondNumber := req.Sum.SecondNumber
	result := fmt.Sprintf("%.2f + %.2f = %.2f", firstNumber, secondNumber, firstNumber+secondNumber)

	res := &calculatorpb.SumResponse{
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
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
