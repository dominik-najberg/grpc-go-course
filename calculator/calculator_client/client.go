package main

import (
	"context"
	"github.com/dominik-najberg/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	log.Println("starting client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	log.Printf("client created: %#v\n", c)

	//doUnary(c)
	doServerStream(c)
}

func doServerStream(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		PrimeNumberDecomposition: &calculatorpb.PrimeNumberDecomposition{
			Number: 120,
		},
	}

	stream, err := c.DecomposePrimeNumber(context.Background(), req)
	if err != nil {
		log.Fatalf("error on DecmposePrimeNumber RPC: %v", err)
	}

	for {
		stream, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("failed during stream: %v", err)
		}

		log.Println(stream.Response)
	}
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.SumRequest{
		Sum: &calculatorpb.Sum{
			FirstNumber:  1900,
			SecondNumber: 77,
		},
	}

	resp, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error on Calculate request: %v", err)
	}

	log.Printf("Calculate result: %v", resp.Result)
}
