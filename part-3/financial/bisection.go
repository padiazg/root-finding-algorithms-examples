package financial

import "math"

// BisectionMethod implements bisection root-finding
type BisectionMethod struct {
	maxIter   int
	threshold float64
}

func NewBisectionMethod(maxIter int, threshold float64) *BisectionMethod {
	return &BisectionMethod{
		maxIter:   maxIter,
		threshold: threshold,
	}
}

// FindRoot implements bisection method
func (s *BisectionMethod) FindRoot(f func(float64) float64, initialGuess float64) (float64, int) {
	a, b := initialGuess-1.0, initialGuess+1.0

	for i := 0; i < s.maxIter; i++ {
		var (
			r  = (a + b) / 2
			fa = f(a)
			fr = f(r)
		)

		if fr == 0.0 || math.Abs(fr) < s.threshold {
			return r, i + 1
		}

		if fa*fr < 0 {
			b = r
		} else {
			a = r
		}
	}

	return (a + b) / 2, s.maxIter
}
