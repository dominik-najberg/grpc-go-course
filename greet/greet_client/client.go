package main

import (
	"context"
	"github.com/dominik-najberg/grpc-go-course/greet/greetpb"
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

	c := greetpb.NewGreetServiceClient(cc)
	log.Printf("client created: %#v\n", c)

	//doUnary(c)
	//doServerStreaming(c)
	doClientStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	log.Println("doing Unary RPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Dominik",
			LastName:  "Najberg",
		},
	}
	resp, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error on Greet RPC: %v", err)
	}

	log.Printf("server response: %s", resp.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	log.Println("doing Server Streaming RPC request")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Dominik",
			LastName:  "Najberg",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error on GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		} else if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}

		log.Println(msg.GetResult())
	}

}

func doClientStreaming(c greetpb.GreetServiceClient) {
	log.Println("doing Client Streaming RPC")

	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Dominik",
				LastName:  "Najberg",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Renata",
				LastName:  "Najberg",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Pola",
				LastName:  "Najberg",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Leon",
				LastName:  "Najberg",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while streaming to server: %v", err)
	}

	for _, request := range requests {
		if err := stream.Send(request); err != nil {
			log.Fatalf("error while sending request: %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp)
}
