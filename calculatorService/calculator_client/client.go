package main

import (
	"context"
	"fmt"
	"log"

	"github.com/odair/gocourse/calculatorService/calculatorpb"
	"google.golang.org/grpc"

)

func main() {
	fmt.Println("Hello I'm calculator client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	log.Println("Starting to do a Unary RPC calculator..")

	req := &calculatorpb.CalculatorRequest{
		Calculating: &calculatorpb.Calculating{
			FirstNumber:  1,
			SecondNumber: 2,
		},
	}

	res, err := c.Calculator(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Calculator RPC: %v", err)
	}

	log.Printf("Result from Calculator: %v", res.Result)
}
