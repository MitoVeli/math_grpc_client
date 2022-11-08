package pkg

type MathGrpcClient interface {
	Calculate(x int64, y int64, operationSign string, result *int64) error
}
