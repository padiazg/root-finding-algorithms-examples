package financial

import (
	"math"
	"testing"
)

// Benchmark BisectionMethod
func BenchmarkBisectionMethod(b *testing.B) {
	cashFlows := createCashFlows()
	method := NewBisectionMethod(1000, 1e-7)
	for i := 0; i < b.N; i++ {
		CalculateIRR(cashFlows, method)
	}
}

// Test BisectionMethod
func TestBisectionMethod(t *testing.T) {
	cashFlows := createCashFlows()
	method := NewBisectionMethod(1000, 1e-7)
	irr, _ := CalculateIRR(cashFlows, method)
	expectedIRR := 0.1634 // Updated expected IRR value

	if math.Abs(irr-expectedIRR) > 1e-5 {
		t.Errorf("BisectionMethod: expected %v, got %v", expectedIRR, irr)
	}
}
