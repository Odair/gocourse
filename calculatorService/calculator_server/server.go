package main

import (
	"context"
	"fmt"
	"log"
	"net"

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
