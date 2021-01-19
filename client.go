package main

import (
	"context"
	"fmt"
	"kubernetes/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Starting client...")

	cc, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v ", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	dounary(c)

}

func dounary(c greetpb.GreetServiceClient) {
	fmt.Println(".......service.......")
	req := &greetpb.GreetingRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Lakshmi",
			LastName:  "M",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling request %v ", err)
	}
	log.Printf(" responce from greet %v ", res)
}
