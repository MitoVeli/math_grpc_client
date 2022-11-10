package math_service

type MathService interface {
	Calculate(x float32, y float32, operationSign string) (float32, error)
}
