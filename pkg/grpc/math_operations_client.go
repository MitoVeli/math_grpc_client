package grpc

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

var (
	defaultRequestTimeout = time.Second * 5
	client                MathOperationsClient
)

// InitializeAuthRpc inits gRPC connection and creates a client
func InitializeMathRpc(connectionString string) {
	InitializeRpc(connectionString, func(c *grpc.ClientConn) {
		client = NewMathOperationsClient(c)
	})
}

// Get user's first and last names by user Id
func Calculate(userId string) (*OperationResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()
	r, err := client.Calculate(ctx, &OperationRequest{X: 1, Y: 2, OperationSign: "+"})

	if err != nil {
		log.Printf("Error while calling Add function: %v", err)
		return nil, err
	}

	return r, err
}
