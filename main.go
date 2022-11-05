package main

import (
	"fmt"
	"log"
	"net/http"

	grpcClient "math_grpc_client/grpc"

	configs "math_grpc_client/configs"

	gppcServer "github.com/MitoVeli/math_grpc_server/pkg"
)

func main() {

	// TODO: Move vars to a config file

	// initialize grpc client
	grpcClient.InitializeMathRpc("localhost:" + configs.GrpcPort)

	// TODO: To be deleted
	newMathOperationsService := gppcServer.NewMathOperationsService()
	var result int32
	err := newMathOperationsService.Add(1, 2, &result)
	if err != nil {
		log.Printf("Error while sending grpc request: %v", err)
	}

	fmt.Println("The result is:", result)

	err = newMathOperationsService.Subtract(1, 2, &result)
	if err != nil {
		log.Printf("Error while sending grpc request: %v", err)
	}

	fmt.Println("The result is:", result)
	//////

	// start http server
	s := http.Server{
		Addr: configs.AppPort,
	}
	s.ListenAndServe()

}
