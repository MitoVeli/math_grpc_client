package math_service

import (
	"log"

	mathClient "github.com/MitoVeli/math_grpc_client/pkg/grpc/math/client"
)

type mathClientService struct {
}

func NewMathClientService() MathService {
	return &mathClientService{}
}

func (s *mathClientService) Calculate(x float32, y float32, operationSign string) (float32, error) {

	result, err := mathClient.Calculate(x, y, operationSign)
	if err != nil {
		log.Fatalf("Error while sending grpc request: %v", err)
		return 0, err
	}

	return result.Result, nil
}
