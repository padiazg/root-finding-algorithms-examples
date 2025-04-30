package financial

import "math"

// SecantMethod implements secant root-finding method
type SecantMethod struct {
	maxIter   int
	threshold float64
}

func NewSecantMethod(maxIter int, threshold float64) *SecantMethod {
	return &SecantMethod{
		maxIter:   maxIter,
		threshold: threshold,
	}
}

// FindRoot implements secant method
func (s *SecantMethod) FindRoot(f func(float64) float64, initialGuess float64) (float64, int) {
	x0 := initialGuess
	x1 := initialGuess + 0.1

	for i := 0; i < s.maxIter; i++ {
		fx0 := f(x0)
		fx1 := f(x1)

		if math.Abs(fx1) < s.threshold {
			return x1, i + 1
		}

		// Secant method formula
		x2 := x1 - fx1*(x1-x0)/(fx1-fx0)

		x0, x1 = x1, x2
	}

	return x1, s.maxIter
}
