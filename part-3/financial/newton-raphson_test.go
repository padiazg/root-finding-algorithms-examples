package financial

import (
	"math"
	"testing"
)

// Test NewtonRaphsonMethod
func TestNewtonRaphsonMethod(t *testing.T) {
	cashFlows := createCashFlows()
	method := NewNewtonRaphsonMethod(1000, 1e-7)
	irr, _ := CalculateIRR(cashFlows, method)
	expectedIRR := 0.1634 // Updated expected IRR value

	if math.Abs(irr-expectedIRR) > 1e-5 {
		t.Errorf("NewtonRaphsonMethod: expected %v, got %v", expectedIRR, irr)
	}
}

// Benchmark NewtonRaphsonMethod
func BenchmarkNewtonRaphsonMethod(b *testing.B) {
	cashFlows := createCashFlows()
	method := NewNewtonRaphsonMethod(1000, 1e-7)
	for i := 0; i < b.N; i++ {
		CalculateIRR(cashFlows, method)
	}
}
