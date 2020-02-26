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
	//doServerStream(c)
	//doClientStream(c)
	doBiDiStream(c)
}

func doBiDiStream(c calculatorpb.CalculatorServiceClient) {
	numbers := []int32{1, 5, 3, 6, 2, 20}
	breakChannel := make(chan bool)

	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// sending data to server
	go func() {
		for _, number := range numbers {
			err := stream.Send(&calculatorpb.FindMaximumRequest{
				Number: number,
			})
			if err != nil {
				breakChannel <- true
				log.Fatal(err)
			}
		}
		err := stream.CloseSend()
		if err != nil {
			breakChannel <- true
			log.Fatal(err)
		}
	}()

	// receiving data from server
	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				breakChannel <- true
				break
			}
			if err != nil {
				breakChannel <- true
				log.Fatal(err)
			}
			log.Println("the maximum value is: ", msg.GetNumber())
		}
	}()

	<-breakChannel
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

func doClientStream(c calculatorpb.CalculatorServiceClient) {
	numbers := []int32{1, 2, 3, 4}

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, number := range numbers {
		err := stream.Send(&calculatorpb.ComputeAverageRequest{
			Number: number,
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	result, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("calculated average: ", result.GetResult())
}
