package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	//doUnary(c)
	//doServerStreaming(c)
	doClienteStreaming(c)
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

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	log.Println("Starting to do a Server Streaming RPC..")

	req := &calculatorpb.CalculatorManyTimesRequest{
		CalculatingManyTimes: &calculatorpb.CalculatingManyTimes{
			Number: 120,
		},
	}

	resStream, err := c.CalculatorManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Calculator stream RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while read stream: %v", err)
		}
		log.Printf("Response from CalculatorManyTimes: %v", msg.GetResult())
	}
}

func doClienteStreaming(c calculatorpb.CalculatorServiceClient) {
	log.Println("Starting to do a Client Streaming Average Numbers..")

	requests := []*calculatorpb.NumberAverageRequest{
		&calculatorpb.NumberAverageRequest{
			NumberAverage: &calculatorpb.NumberAverage{
				Number: 1,
			},
		},
		&calculatorpb.NumberAverageRequest{
			NumberAverage: &calculatorpb.NumberAverage{
				Number: 2,
			},
		},
		&calculatorpb.NumberAverageRequest{
			NumberAverage: &calculatorpb.NumberAverage{
				Number: 3,
			},
		},
		&calculatorpb.NumberAverageRequest{
			NumberAverage: &calculatorpb.NumberAverage{
				Number: 4,
			},
		},
	}

	stream, err := c.NumberAverage(context.Background())
	if err != nil {
		log.Fatalf("error while calling NumberAverage: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from NumberAverage: %v", err)
	}

	fmt.Printf("NumberAverage Response: %v\n", res.Result)
}
