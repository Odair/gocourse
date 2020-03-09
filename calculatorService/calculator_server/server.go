package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/odair/gocourse/calculatorService/calculatorpb"
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
	number := req.GetCalculatingManyTimes().GetNumber()
	var k int32 = 2

	for number > 1 {
		if number%k == 0 {
			result := k
			res := &calculatorpb.CalculatorManyTimesResponse{
				Result: result,
			}
			stream.Send(res)
			time.Sleep(1000 * time.Millisecond)
			number = number / k
		} else {
			k++
			fmt.Println("New divisor is: &v", k)
		}
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
