package main

import (
	"flag"
	"log"

	configs "github.com/MitoVeli/math_grpc_client/configs"
	grpcClient "github.com/MitoVeli/math_grpc_client/grpc"
	grpcService "github.com/MitoVeli/math_grpc_client/pkg"
	gppcServer "github.com/MitoVeli/math_grpc_server/pkg"
)

func main() {

	// TODO: tests
	// TODO: check config
	// TODO: check project structure, server, docker etc
	// TODO: check folder structure as asked!
	// TODO: check all comments and add if necessary
	// TODO: consider changing to float64
	// TODO: check imports within project

	// declare flags
	var firstNum int64
	var secondNum int64
	var operationSign string

	// receive and parse flags
	flag.Int64Var(&firstNum, "firstNum", 0, "first number")
	flag.Int64Var(&secondNum, "secondNum", 0, "second number")
	flag.StringVar(&operationSign, "operationSign", "", "operation sign")
	flag.Parse()

	// initialize grpc client
	grpcClient.InitializeMathRpc("localhost:" + configs.GrpcPort)

	// initialize new math operation server
	mathGrpcServer := gppcServer.NewMathOperationsService()

	// initialize new math grpc client
	mathGrpcClient := grpcService.NewMathClientService(mathGrpcServer)

	// trigger client math operation with given parameters
	err := mathGrpcClient.Calculate(firstNum, secondNum, operationSign)
	if err != nil {
		log.Fatalf("Error while sending grpc request: %v", err)
		return
	}

}
