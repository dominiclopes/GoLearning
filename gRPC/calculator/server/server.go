package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"calculator/calculatorpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (* server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (resp *calculatorpb.SumResponse, err error) {
	fmt.Println("Executing Sum function via GRPC")

	// get request params
	firtNumber := req.GetFirstNumber()
	secondNumber := req.GetSecondNumber()

	// execute the business logic
	result := firtNumber + secondNumber

	// prepare the response
	resp = &calculatorpb.SumResponse{
		SumResult: result,
	}

	// return the response
	return resp, nil
}


func main() {
	fmt.Println("Starting thr gRPC server....")

	// Create a connection
	listen, err := net.Listen("tcp", "0.0.0.0:5051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a server
	s := grpc.NewServer()

	// Register the service on server
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	// Register reflection service on gRPC service
	reflection.Register(s)

	// Start listening on the created connection
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
