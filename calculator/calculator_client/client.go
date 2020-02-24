package main

import (
	"context"
	"github.com/dominik-najberg/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
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

	doUnary(c)
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
