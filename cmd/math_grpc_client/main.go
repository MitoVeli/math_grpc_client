package main

import (
	"flag"
	"log"
	"net/http"

	configs "github.com/MitoVeli/math_grpc_client/configs"
	grpcClient "github.com/MitoVeli/math_grpc_client/pkg/grpc/math/client"
	mathService "github.com/MitoVeli/math_grpc_client/pkg/math_service"
)

func main() {

	// declare flags
	var firstNum float64
	var secondNum float64
	var operationSign string

	// receive and parse flags
	flag.Float64Var(&firstNum, "firstNum", 0.0, "first number")
	flag.Float64Var(&secondNum, "secondNum", 0.0, "second number")
	flag.StringVar(&operationSign, "operationSign", "", "operation sign")
	flag.Parse()

	// convert firstNum from float64 to float32
	x := float32(firstNum)
	y := float32(secondNum)

	// check app environment and set connection string for grpc server
	environment := configs.SetEnv()

	// initialize grpc client
	grpcClient.InitializeMathRpc(environment + configs.GrpcPort)

	newMathService := mathService.NewMathClientService()

	result, err := newMathService.Calculate(x, y, operationSign)
	if err != nil {
		log.Fatalf("Error while sending grpc request: %v", err)
		return
	}

	log.Printf("Result is: %v", result)

	// start http server
	s := http.Server{
		Addr: configs.AppPort,
	}
	s.ListenAndServe()

}
