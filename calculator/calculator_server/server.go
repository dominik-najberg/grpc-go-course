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

func (s *server) DecomposePrimeNumber(
	req *calculatorpb.PrimeNumberDecompositionRequest,
	stream calculatorpb.CalculatorService_DecomposePrimeNumberServer,
) error {
	log.Printf("DecomposePrimeNumber function invoked with: %v", req)
	number := int(req.GetPrimeNumberDecomposition().GetNumber())

	k := 2

	for number > 1 {
		if number%k == 0 {
			resp := &calculatorpb.PrimeNumberDecompositionResponse{
				Response: int32(k),
			}
			stream.Send(resp)
			log.Println(k)
			number = number / k // divide N by k so that we have the rest of the number left.
		} else { // if k evenly divides into N
			k = k + 1
		}
	}

	return nil
}

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
