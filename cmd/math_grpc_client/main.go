package main

import (
	"flag"
	"log"
	"net/http"

	grpcClient "github.com/MitoVeli/math_grpc_client/pkg/grpc"
	grpcService "github.com/MitoVeli/math_grpc_client/pkg/math_client"
	grpcServer "github.com/MitoVeli/math_grpc_server/pkg/math"
)

func main() {

	// TODO: check folder structure as asked!
	// TODO: consider changing to float64

	// declare flags
	var firstNum float64
	var secondNum float64
	var operationSign string

	// receive and parse flags
	flag.Float64Var(&firstNum, "firstNum", 0.0, "first number")
	flag.Float64Var(&secondNum, "secondNum", 0.0, "second number")
	flag.StringVar(&operationSign, "operationSign", "", "operation sign")
	flag.Parse()

	if operationSign == "" {
		operationSign = "+"
	}

	// convert firstNum from float64 to float32
	x := float32(firstNum)
	y := float32(secondNum)

	// initialize new math operation server
	mathGrpcServer := grpcServer.NewMathOperationsService()

	// initialize grpc client
	grpcClient.InitializeMathRpc("localhost: 50051")

	// initialize new math grpc client
	mathGrpcClient := grpcService.NewMathClientService(mathGrpcServer)

	// trigger client math operation with given parameters
	var result float32
	err := mathGrpcClient.Calculate(x, y, operationSign, &result)
	if err != nil {
		log.Fatalf("Error while sending grpc request: %v", err)
		return
	}

	log.Printf("Result is: %v", float64(result))

	// start http server
	s := http.Server{
		Addr: ":8008",
	}
	s.ListenAndServe()

}
