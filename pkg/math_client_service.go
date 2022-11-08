package pkg

import (
	"log"

	mathGrpcServer "github.com/MitoVeli/math_grpc_server/pkg"
)

type mathClientService struct {
	mathServer mathGrpcServer.MathOperations
}

func NewMathClientService(mathGrpcServer mathGrpcServer.MathOperations) MathGrpcClient {
	return &mathClientService{
		mathServer: mathGrpcServer,
	}
}

func (s *mathClientService) Calculate(x int64, y int64, operationSign string) error {

	var result int64
	// call math operation grpc server
	err := s.mathServer.DoMath(x, y, operationSign, &result)
	if err != nil {
		log.Fatalf("Error while sending grpc request: %v", err)
		return err
	}

	log.Printf("Result is: %v", result)
	return nil
}
