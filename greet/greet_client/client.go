package main

import (
	"context"
	"github.com/dominik-najberg/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"time"
)

func main() {
	log.Println("starting client")

	creds, sslErr := credentials.NewClientTLSFromFile("ssl/ca.crt", "")
	if sslErr != nil {
		log.Fatalf("failed while loading certificates: %v", sslErr)
	}

	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	log.Printf("client created: %#v\n", c)

	//doUnary(c)
	//doServerStreaming(c)
	//doClientStreaming(c)
	//doBiDiStreaming(c)
	doUnaryWithDeadline(c, 4*time.Second) // should complete
	doUnaryWithDeadline(c, 1*time.Second) // should timeout
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

	requests := createRequests()

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

func createRequests() []*greetpb.LongGreetRequest {
	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Dominik",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Renata",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Pola",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Leon",
			},
		},
	}
	return requests
}

func doBiDiStreaming(c greetpb.GreetServiceClient) {
	log.Println("doing BiDi Client Streaming RPC")

	requests := createRequests()

	// create a stream by invoking a client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	waitChannel := make(chan bool)

	// we send a bunch of messages to the client (go routine)
	go func() {
		for _, request := range requests {
			err := stream.Send(&greetpb.GreetEveryoneRequest{
				Greeting: request.GetGreeting(),
			})
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
		err := stream.CloseSend()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// we receive messages from the client (go routine)
	go func() {
		for {
			messages, err := stream.Recv()
			if err == io.EOF {
				waitChannel <- true
				break
			}
			if err != nil {
				log.Fatalf("error on stream receive: %v", err)
			}

			log.Println(messages.Result)
		}
	}()

	// block until we're done
	<-waitChannel
}

func doUnaryWithDeadline(c greetpb.GreetServiceClient, s time.Duration) {
	log.Println("doing UnaryWithDeadline RPC and timeout set to:", s)
	req := &greetpb.GreetWithDeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Dominik",
			LastName:  "Najberg",
		},
	}

	d := time.Now().Add(s)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	resp, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		if statusError, ok := status.FromError(err); ok {
			if statusError.Code() == codes.DeadlineExceeded {
				log.Fatalf("request timeout: %v", statusError.Message())
			} else {
				log.Fatalf("unexpected error: %v", statusError.Message())
			}
		} else {
			log.Fatalf("error on GreetWithDeadline RPC: %v", err)
		}
	}

	log.Printf("server response: %s", resp.Result)
}
