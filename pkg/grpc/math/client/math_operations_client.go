package client

import (
	context "context"
	"log"
	"time"

	"github.com/MitoVeli/math_grpc_client/pkg/grpc/factory"
	pb "github.com/MitoVeli/math_grpc_client/pkg/grpc/math"

	"google.golang.org/grpc"
)

var (
	defaultRequestTimeout = time.Second * 5
	client                pb.MathOperationsClient
)

// InitializeAuthRpc inits gRPC connection and creates a client
func InitializeMathRpc(connectionString string) {
	factory.InitializeRpc(connectionString, func(c *grpc.ClientConn) {
		client = pb.NewMathOperationsClient(c)
	})
}

func Calculate(x float32, y float32, operationSign string) (*pb.OperationResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	// call math operation grpc server
	res, err := client.Calculate(ctx, &pb.OperationRequest{
		X:             x,
		Y:             y,
		OperationSign: operationSign,
	})
	if err != nil {
		log.Fatalf("Error while sending grpc request: %v", err)
		return nil, err
	}

	return res, nil
}
