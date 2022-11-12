package client

import (
	context "context"
	"log"
	"time"

	"github.com/MitoVeli/math_grpc_client/pkg/grpc/factory"
	pb "github.com/MitoVeli/math_grpc_client/pkg/grpc/math"

	"google.golang.org/grpc"
)

type mathGrpcClient struct {
	defaultRequestTimeout time.Duration
	client                pb.MathOperationsClient
}

func NewMathGrpcClient(connectionString string) MathGrpcClient {
	var client pb.MathOperationsClient
	factory.InitializeRpc(connectionString, func(c *grpc.ClientConn) {
		client = pb.NewMathOperationsClient(c)
	})
	return &mathGrpcClient{
		defaultRequestTimeout: time.Second * 5,
		client:                client,
	}
}

func (c mathGrpcClient) Calculate(x float32, y float32, operationSign string) (*pb.OperationResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.defaultRequestTimeout)
	defer cancel()

	// call math operation grpc server
	res, err := c.client.Calculate(ctx, &pb.OperationRequest{
		X:             x,
		Y:             y,
		OperationSign: operationSign,
	})

	if err != nil {
		log.Fatalf("Error while sending grpc request: %v", err)
		return nil, err
	}

	return res, err

}
