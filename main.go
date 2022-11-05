package main

import (
	"fmt"
	"net/http"

	grpcClient "github.com/MitoVeli/math_grpc_client/grpc"
)

func main() {

	// TODO: Move vars to a config file

	/* ------------- INITIALIZE gRPC CONNECTIONS ------------- */
	grpcClient.InitializeMathRpc("localhost:50051")

	// start http server
	s := http.Server{
		Addr: ":8008",
	}
	s.ListenAndServe()

	fmt.Println("serving at 8008")

}
