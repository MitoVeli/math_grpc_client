package client

import pb "github.com/MitoVeli/math_grpc_client/pkg/grpc/math"

type MathGrpcClient interface {
	Calculate(x float32, y float32, operationSign string) (*pb.OperationResponse, error)
}
