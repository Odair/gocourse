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
	//doClienteStreaming(c)
	doBidiStreaming(c)
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

func doBidiStreaming(c calculatorpb.CalculatorServiceClient) {
	log.Println("Starting to do a BiDi Streaming RPC..")

	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongGreet: %v", err)
	}

	numbers := []int32{1, 2, 6, 1, 20}
	waitc := make(chan struct{})

	go func() {

		for _, number := range numbers {
			fmt.Printf("Sending request %v\n", number)
			stream.Send(&calculatorpb.FindMaximumRequest{
				Number: number,
			})
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("error while receiving response from FindMaximum: %v", err)
				break
			}

			fmt.Printf("Received new maximum of... %v\n", res.GetMaximum())
		}
		close(waitc)
	}()

	<-waitc
}
