package main

import (
	"flag"
	"fmt"
	"log"

	grpcClient "github.com/MitoVeli/math_grpc_client/grpc"

	configs "github.com/MitoVeli/math_grpc_client/configs"
	gppcServer "github.com/MitoVeli/math_grpc_server/pkg"
)

func main() {

	// declare flags
	var firstNum int64
	var secondNum int64
	var operationSign string

	// receive and parse flags
	flag.Int64Var(&firstNum, "firstNum", 0, "first number")
	flag.Int64Var(&secondNum, "secondNum", 0, "second number")
	flag.StringVar(&operationSign, "operationSign", "", "operation sign")
	flag.Parse()

	log.Println("firstNum:", firstNum, "secondNum:", secondNum, "operationSign:", operationSign)

	// initialize grpc client
	grpcClient.InitializeMathRpc("localhost:" + configs.GrpcPort)

	// initialize new math operation server
	newMathOperationsService := gppcServer.NewMathOperationsService()
	var result int64
	err := newMathOperationsService.DoMath(firstNum, secondNum, operationSign, &result)
	if err != nil {
		log.Printf("Error while sending grpc request: %v", err)
		return
	}

	fmt.Println("The result is:", result)

}
