package math_service

import (
	"log"

	grpcClient "github.com/MitoVeli/math_grpc_client/pkg/grpc/math/client"
)

type mathClientService struct {
	client grpcClient.MathGrpcClient
}

func NewMathClientService(client grpcClient.MathGrpcClient) MathService {
	return &mathClientService{
		client: client,
	}
}

func (s *mathClientService) DoMath(x float32, y float32, operationSign string) (float32, error) {

	result, err := s.client.Calculate(x, y, operationSign)
	if err != nil {
		log.Printf("Error while sending grpc request: %v", err)
		return 0, err
	}

	return result.Result, nil
}
