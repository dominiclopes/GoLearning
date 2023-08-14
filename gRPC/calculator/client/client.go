package main

import (
	"calculator/calculatorpb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)


func main(){
	// create a client connection
	cc, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to created client stub: %v", err)
	}
	defer cc.Close()

	// create
	conn := calculatorpb.NewCalculatorServiceClient(cc)

	sumNumbers(conn)
}


func sumNumbers(conn calculatorpb.CalculatorServiceClient) {
	fmt.Println("Executing sum of number via gRPC function")

	// prepare the request
	firstNumber := 3
	secondNumber := 5
	req := &calculatorpb.SumRequest{
		FirstNumber: int32(firstNumber),
		SecondNumber: int32(secondNumber),
	}

	// execute the rpc function
	resp, err := conn.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error serving the request: %v", err)
	}

	// display the response
	fmt.Println("The sum of number is", resp)
}