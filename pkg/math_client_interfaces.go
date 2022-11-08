package pkg

type MathGrpcClient interface {
	Calculate(x int64, y int64, operationSign string) error
}
