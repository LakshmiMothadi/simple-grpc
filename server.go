package main

import (
	"context"
	"fmt"
	"kubernetes/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetingRequest) (*greetpb.GreetingResponse, error) {
	fmt.Printf("greet function is invoked with %v ", req)
	FirstName := req.GetGreeting().GetFirstName()
	result := "hello" + FirstName
	res := &greetpb.GreetingResponse{
		Result: result,
	}
	fmt.Println("hello from server")
	return res, nil
}
func main() {
	fmt.Println("hello")
	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("error while calling %v ", err)

	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
