package main

import (
	"flag"
	"log"

	grpcClient "github.com/MitoVeli/math_grpc_client/grpc"

	configs "github.com/MitoVeli/math_grpc_client/configs"
	gppcServer "github.com/MitoVeli/math_grpc_server/pkg"
)

func main() {

	// TODO: tests
	// TODO: check config
	// TODO: check project structure, server, docker etc
	// TODO: check folder structure as asked!
	// TODO: check all comments and add if necessary

	// declare flags
	var firstNum int64
	var secondNum int64
	var operationSign string

	// receive and parse flags
	flag.Int64Var(&firstNum, "firstNum", 0, "first number")
	flag.Int64Var(&secondNum, "secondNum", 0, "second number")
	flag.StringVar(&operationSign, "operationSign", "", "operation sign")
	flag.Parse()

	log.Printf("firstNum: %d, secondNum: %d, operationSign: %s", firstNum, secondNum, operationSign)

	// TODO: check below part again!!!
	// initialize grpc client
	grpcClient.InitializeMathRpc("localhost:" + configs.GrpcPort)

	// initialize new math operation server
	newMathOperationsService := gppcServer.NewMathOperationsService()

	var result int64
	// call math operation grpc server
	err := newMathOperationsService.DoMath(firstNum, secondNum, operationSign, &result)
	if err != nil {
		log.Fatalf("Error while sending grpc request: %v", err)
		return
	}

	log.Printf("Result is: %v", result)

}
