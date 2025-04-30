package financial

import "math"

// CalculateIRR computes Internal Rate of Return using provided root-finding method
func CalculateIRR(cashFlows []CashFlow, method RootFinder) (float64, int) {
	// Create NPV function for IRR calculation
	npvFunc := func(rate float64) float64 {
		npv := 0.0
		for _, cf := range cashFlows {
			npv += cf.Amount / math.Pow(1+rate, float64(cf.Period))
		}
		return npv
	}

	// Initial guess and root finding
	root, iterations := method.FindRoot(npvFunc, 0.1)
	return root, iterations
}
