package test

import (
	"errors"
	"testing"

	"github.com/MitoVeli/math_grpc_client/pkg/grpc/math"
	mockGrpcClient "github.com/MitoVeli/math_grpc_client/pkg/grpc/math/client/mocks"
	mathService "github.com/MitoVeli/math_grpc_client/pkg/math_service"
	"github.com/stretchr/testify/assert"
)

// Math Service Client Tests
func TestDoMath(t *testing.T) {

	firstNumber := float32(10)
	secondNumber := float32(5)
	invalidOperationSign := "&"

	mockMathClient := new(mockGrpcClient.MathGrpcClient)
	mathClientService := mathService.NewMathClientService(mockMathClient)

	t.Run("Add", func(t *testing.T) {

		mockMathClient.On("Calculate", firstNumber, secondNumber, "+").Return(&math.OperationResponse{
			Result: firstNumber + secondNumber,
		}, nil)

		// trigger math operation service with given parameters
		res, err := mathClientService.DoMath(firstNumber, secondNumber, "+")

		assert.NoError(t, err)
		assert.Equal(t, firstNumber+secondNumber, res)
	})

	t.Run("Subtract", func(t *testing.T) {

		mockMathClient.On("Calculate", firstNumber, secondNumber, "-").Return(&math.OperationResponse{
			Result: firstNumber - secondNumber,
		}, nil)

		// trigger math operation service with given parameters
		res, err := mathClientService.DoMath(firstNumber, secondNumber, "-")

		assert.NoError(t, err)
		assert.Equal(t, firstNumber-secondNumber, res)
	})

	t.Run("Multiply", func(t *testing.T) {

		mockMathClient.On("Calculate", firstNumber, secondNumber, "*").Return(&math.OperationResponse{
			Result: firstNumber * secondNumber,
		}, nil)

		// trigger math operation service with given parameters
		res, err := mathClientService.DoMath(firstNumber, secondNumber, "*")

		assert.NoError(t, err)
		assert.Equal(t, firstNumber*secondNumber, res)
	})

	t.Run("Divide", func(t *testing.T) {

		mockMathClient.On("Calculate", firstNumber, secondNumber, "/").Return(&math.OperationResponse{
			Result: firstNumber / secondNumber,
		}, nil)

		// trigger math operation service with given parameters
		res, err := mathClientService.DoMath(firstNumber, secondNumber, "/")

		assert.NoError(t, err)
		assert.Equal(t, firstNumber/secondNumber, res)
	})

	t.Run("Divide by zero", func(t *testing.T) {

		mockMathClient.On("Calculate", float32(0), float32(0), "/").Return(nil, errors.New("Error while sending grpc request: cannot divide by zero"))

		// trigger math operation service with given parameters
		_, err := mathClientService.DoMath(float32(0), float32(0), "/")

		assert.Error(t, err, "Error while sending grpc request: cannot divide by zero")
	})

	t.Run("Invalid operation sign", func(t *testing.T) {

		mockMathClient.On("Calculate", firstNumber, secondNumber, invalidOperationSign).Return(
			nil, errors.New("Error while sending grpc request: invalid operation sign"))

		// trigger math operation service with given parameters
		_, err := mathClientService.DoMath(firstNumber, secondNumber, invalidOperationSign)

		assert.Error(t, err, "Error while sending grpc request: invalid operation sign")
	})
}
