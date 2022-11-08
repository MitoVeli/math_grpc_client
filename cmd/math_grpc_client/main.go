package main

import (
	"flag"
	"log"
	"net/http"

	grpcClient "github.com/MitoVeli/math_grpc_client/pkg/grpc"
	grpcService "github.com/MitoVeli/math_grpc_client/pkg/math_client"
	gppcServer "github.com/MitoVeli/math_grpc_server/pkg"
)

func main() {

	// TODO: tests
	// TODO: check config
	// TODO: check project structure, server, docker etc
	// TODO: check folder structure as asked!
	// TODO: check all comments and add if necessary
	// TODO: consider changing to float64

	// declare flags
	var firstNum int64
	var secondNum int64
	var operationSign string

	// receive and parse flags
	flag.Int64Var(&firstNum, "firstNum", 0, "first number")
	flag.Int64Var(&secondNum, "secondNum", 0, "second number")
	flag.StringVar(&operationSign, "operationSign", "", "operation sign")
	flag.Parse()

	if operationSign == "" {
		operationSign = "+"
	}

	// initialize new math operation server
	mathGrpcServer := gppcServer.NewMathOperationsService()

	// initialize grpc client
	grpcClient.InitializeMathRpc("localhost: 50051")

	// initialize new math grpc client
	mathGrpcClient := grpcService.NewMathClientService(mathGrpcServer)

	// trigger client math operation with given parameters
	err := mathGrpcClient.Calculate(firstNum, secondNum, operationSign)
	if err != nil {
		log.Fatalf("Error while sending grpc request: %v", err)
		return
	}

	// start http server
	s := http.Server{
		Addr: ":8008",
	}
	s.ListenAndServe()

}
