package test

import (
	"testing"

	grpcService "github.com/MitoVeli/math_grpc_client/pkg/math_client"
	mockMathGrpcServer "github.com/MitoVeli/math_grpc_server/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {

	firstNumber := int64(10)
	secondNumber := int64(5)
	var result int64

	mockMathServer := new(mockMathGrpcServer.MathOperations)

	mathGrpcClient := grpcService.NewMathClientService(mockMathServer)

	t.Run("Add", func(t *testing.T) {

		mockMathServer.On("DoMath", firstNumber, secondNumber, "+", &result).Return(nil)

		// trigger client math operation with given parameters
		err := mathGrpcClient.Calculate(firstNumber, secondNumber, "+")

		assert.NoError(t, err)
	})

	t.Run("Subtract", func(t *testing.T) {

		mockMathServer.On("DoMath", firstNumber, secondNumber, "-", &result).Return(nil)

		// trigger client math operation with given parameters
		err := mathGrpcClient.Calculate(firstNumber, secondNumber, "-")

		assert.NoError(t, err)
	})

	t.Run("Multiply", func(t *testing.T) {

		mockMathServer.On("DoMath", firstNumber, secondNumber, "*", &result).Return(nil)

		// trigger client math operation with given parameters
		err := mathGrpcClient.Calculate(firstNumber, secondNumber, "*")

		assert.NoError(t, err)
	})

	t.Run("Divide", func(t *testing.T) {

		mockMathServer.On("DoMath", firstNumber, secondNumber, "/", &result).Return(nil)

		// trigger client math operation with given parameters
		err := mathGrpcClient.Calculate(firstNumber, secondNumber, "/")

		assert.NoError(t, err)
	})
}
