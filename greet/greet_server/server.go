package main

import (
	"context"
	"fmt"
	"github.com/dominik-najberg/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"time"
)

type server struct{}

func (s *server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {
	log.Printf("Greet function invoked with: %v", req)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			log.Println("the client canceled the request")
			return nil, status.Error(codes.Canceled, "client canceled")
		}
		time.Sleep(time.Second)
	}

	firstName := req.GetGreeting().FirstName
	result := fmt.Sprintf("Hello, %s", firstName)

	res := &greetpb.GreetWithDeadlineResponse{
		Result: result,
	}

	return res, nil
}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greet function invoked with: %v", req)

	firstName := req.GetGreeting().FirstName
	result := fmt.Sprintf("Hello, %s", firstName)

	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func (s *server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function invoked with: %v", req)

	firstName := req.GetGreeting().GetFirstName()

	for i := 0; i < 10; i++ {
		result := fmt.Sprintf("Hello, %s! Times: %d", firstName, i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		if err := stream.Send(res); err != nil {
			return err
		}

		time.Sleep(time.Second)
	}

	return nil
}

func (s *server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function invoked with a streaming request")
	result := "Hello "

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// finished reading client stream
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		} else if err != nil {
			return err
		}

		result += req.GetGreeting().GetFirstName() + "! "
	}
}

func (s *server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone function invoked with a streaming request")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := fmt.Sprintf("Hello, %s", firstName)
		err = stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if err != nil {
			return err
		}
	}
}

func main() {
	log.Println("server running...")

	creds, sslErr := credentials.NewServerTLSFromFile("ssl/server.crt", "ssl/server.pem")
	if sslErr != nil {
		log.Fatalf("failed while loading certificates: %v", sslErr)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
