package pkg

type MathGrpcClient interface {
	Calculate(x float32, y float32, operationSign string, result *float32) error
}
