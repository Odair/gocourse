package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/odair/gocourse/calculatorService/calculatorpb"
	"github.com/odair/gocourse/greet/greetpb"
	"google.golang.org/grpc"

)

type server struct{}

func (*server) Calculator(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	log.Printf("Calculator function was invoked by: %v", req)
	firstNumber := req.GetCalculating().GetFirstNumber()
	secondNumber := req.GetCalculating().GetSecondNumber()

	result := (firstNumber + secondNumber)

	res := &calculatorpb.CalculatorResponse{
		Result: result,
	}
	return res, nil
}

func (*server) CalculatorManyTimes(req *calculatorpb.CalculatorManyTimesRequest, stream calculatorpb.CalculatorService_CalculatorManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked by: %v", req)
	number := req.GetCalculating().GetNumber()

	for i := 0; i <= 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func main() {
	fmt.Println("Hello Calculator")

	listen, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
