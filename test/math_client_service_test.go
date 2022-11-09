package test

import (
	"testing"

	grpcService "github.com/MitoVeli/math_grpc_client/pkg/math_client"
	mockMathGrpcServer "github.com/MitoVeli/math_grpc_server/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {

	firstNumber := float32(10)
	secondNumber := float32(5)
	var result float32

	mockMathServer := new(mockMathGrpcServer.MathOperations)

	mathGrpcClient := grpcService.NewMathClientService(mockMathServer)

	t.Run("Add", func(t *testing.T) {

		mockMathServer.On("DoMath", firstNumber, secondNumber, "+", &result).Return(nil)

		// trigger client math operation with given parameters
		err := mathGrpcClient.Calculate(firstNumber, secondNumber, "+", &result)

		assert.NoError(t, err)
	})

	t.Run("Subtract", func(t *testing.T) {

		mockMathServer.On("DoMath", firstNumber, secondNumber, "-", &result).Return(nil)

		// trigger client math operation with given parameters
		err := mathGrpcClient.Calculate(firstNumber, secondNumber, "-", &result)

		assert.NoError(t, err)
	})

	t.Run("Multiply", func(t *testing.T) {

		mockMathServer.On("DoMath", firstNumber, secondNumber, "*", &result).Return(nil)

		// trigger client math operation with given parameters
		err := mathGrpcClient.Calculate(firstNumber, secondNumber, "*", &result)

		assert.NoError(t, err)
	})

	t.Run("Divide", func(t *testing.T) {

		mockMathServer.On("DoMath", firstNumber, secondNumber, "/", &result).Return(nil)

		// trigger client math operation with given parameters
		err := mathGrpcClient.Calculate(firstNumber, secondNumber, "/", &result)

		assert.NoError(t, err)
	})
}
