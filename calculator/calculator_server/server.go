package main

import (
	"context"
	"fmt"
	"github.com/dominik-najberg/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"math"
	"net"
)

type server struct{}

func (s *server) SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
	log.Printf("SquareRoot function invoked with: %v", req)

	number := req.GetNumber()

	if number < 0 {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("input cannot be negative: %v", number))
	}

	resp := &calculatorpb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}

	return resp, nil
}

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

func (s *server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	log.Printf("ComputeAverage function invoked with stream request")

	var numbers []int32

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Result: calculateAverage(numbers),
			})
		}
		if err != nil {
			log.Fatalf("error while receiving data from client: %v", err)
		}
		numbers = append(numbers, msg.GetNumber())
		log.Println("number received: ", msg.GetNumber())
	}
}

func (s *server) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {
	log.Printf("FindMaximum function invoked with stream request")

	var maxValue int32

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if msg.GetNumber() > maxValue {
			maxValue = msg.GetNumber()
			err = stream.Send(&calculatorpb.FindMaximumResponse{
				Number: maxValue,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	log.Println("server running...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051") // default port for GRPC
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	reflection.Register(s) // Evans GRPC Client - universal client for browsing GRPC servers

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func calculateAverage(args []int32) float64 {
	var sum int32

	for _, arg := range args {
		sum += arg
	}

	return float64(sum) / float64(len(args))
}
