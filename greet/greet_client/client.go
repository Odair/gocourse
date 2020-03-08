package main

import (
	"context"
	"fmt"
	"log"

	"github.com/odair/gocourse/greet/greetpb"
	"google.golang.org/grpc"

)

func main() {
	fmt.Println("Hello I'm client")

	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	log.Println("Starting to do a Unary RPC..")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Odair",
			LastName:  "Santiago",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}
