package financial

import "math"

// NewtonRaphsonMethod implements Newton-Raphson root-finding
type NewtonRaphsonMethod struct {
	maxIter   int
	threshold float64
}

func NewNewtonRaphsonMethod(maxIter int, threshold float64) *NewtonRaphsonMethod {
	return &NewtonRaphsonMethod{
		maxIter:   maxIter,
		threshold: threshold,
	}
}

// FindRoot implements Newton-Raphson method
func (nr *NewtonRaphsonMethod) FindRoot(f func(float64) float64, initialGuess float64) (float64, int) {
	x := initialGuess
	h := 1e-7 // Small value for numerical derivative

	for i := 0; i < nr.maxIter; i++ {
		fx := f(x)
		if math.Abs(fx) < nr.threshold {
			return x, i + 1
		}

		// Numerical derivative approximation
		dfx := (f(x+h) - fx) / h
		if dfx == 0 {
			break
		}

		x = x - fx/dfx
	}

	return x, nr.maxIter
}
