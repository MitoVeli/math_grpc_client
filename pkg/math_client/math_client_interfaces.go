package pkg

type MathGrpcClient interface {
	Calculate(x float64, y float64, operationSign string, result *float64) error
}
