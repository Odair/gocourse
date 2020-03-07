package main

import (

	"fmt"
	"log"

	"github.com/odair/gocourse/greet/greetpb"
	"google.golang.org/grpc"
)

func main(){
	fmt.Println("Hello I'm client")

	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Println("Created client: %f", c)
}