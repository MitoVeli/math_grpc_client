package test

// import (
// 	"errors"
// 	"testing"

// 	grpcService "github.com/MitoVeli/math_grpc_client/pkg/grpc/math/math_client"
// 	mockMathGrpcServer "github.com/MitoVeli/math_grpc_server/pkg/mocks"
// 	"github.com/stretchr/testify/assert"
// )

// func TestCalculate(t *testing.T) {

// 	firstNumber := float32(10)
// 	secondNumber := float32(5)
// 	invalidOperationSign := "&"
// 	var result float32

// 	mockMathServer := new(mockMathGrpcServer.MathOperations)

// 	mathGrpcClient := grpcService.NewMathClientService(mockMathServer)

// 	t.Run("Add", func(t *testing.T) {

// 		mockMathServer.On("DoMath", firstNumber, secondNumber, "+", &result).Return(nil)

// 		// trigger client math operation with given parameters
// 		err := mathGrpcClient.Calculate(firstNumber, secondNumber, "+", &result)

// 		assert.NoError(t, err)
// 	})

// 	t.Run("Subtract", func(t *testing.T) {

// 		mockMathServer.On("DoMath", firstNumber, secondNumber, "-", &result).Return(nil)

// 		// trigger client math operation with given parameters
// 		err := mathGrpcClient.Calculate(firstNumber, secondNumber, "-", &result)

// 		assert.NoError(t, err)
// 	})

// 	t.Run("Multiply", func(t *testing.T) {

// 		mockMathServer.On("DoMath", firstNumber, secondNumber, "*", &result).Return(nil)

// 		// trigger client math operation with given parameters
// 		err := mathGrpcClient.Calculate(firstNumber, secondNumber, "*", &result)

// 		assert.NoError(t, err)
// 	})

// 	t.Run("Divide", func(t *testing.T) {

// 		mockMathServer.On("DoMath", firstNumber, secondNumber, "/", &result).Return(nil)

// 		// trigger client math operation with given parameters
// 		err := mathGrpcClient.Calculate(firstNumber, secondNumber, "/", &result)

// 		assert.NoError(t, err)
// 	})

// 	t.Run("Divide by zero", func(t *testing.T) {

// 		mockMathServer.On("DoMath", float32(0), float32(0), "/", &result).Return(errors.New("Error while sending grpc request: cannot divide by zero"))

// 		// trigger client math operation with given parameters
// 		err := mathGrpcClient.Calculate(float32(0), float32(0), "/", &result)

// 		assert.EqualError(t, err, "Error while sending grpc request: cannot divide by zero")
// 	})

// 	t.Run("Divide by zero", func(t *testing.T) {

// 		mockMathServer.On("DoMath", float32(0), float32(0), "/", &result).Return(errors.New("Error while sending grpc request: cannot divide by zero"))

// 		// trigger client math operation with given parameters
// 		err := mathGrpcClient.Calculate(float32(0), float32(0), "/", &result)

// 		assert.EqualError(t, err, "Error while sending grpc request: cannot divide by zero")
// 	})

// 	t.Run("Invalid operation sign", func(t *testing.T) {

// 		mockMathServer.On("DoMath", firstNumber, secondNumber, invalidOperationSign, &result).Return(errors.New("Error while sending grpc request: invalid operation sign"))

// 		// trigger client math operation with given parameters
// 		err := mathGrpcClient.Calculate(firstNumber, secondNumber, invalidOperationSign, &result)

// 		assert.EqualError(t, err, "Error while sending grpc request: invalid operation sign")
// 	})
// }
