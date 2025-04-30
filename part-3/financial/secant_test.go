package financial

import (
	"math"
	"testing"
)

// Test SecantMethod
func TestSecantMethod(t *testing.T) {
	cashFlows := createCashFlows()
	method := NewSecantMethod(1000, 1e-7)
	irr, _ := CalculateIRR(cashFlows, method)
	expectedIRR := 0.1634 // Updated expected IRR value

	if math.Abs(irr-expectedIRR) > 1e-5 {
		t.Errorf("SecantMethod: expected %v, got %v", expectedIRR, irr)
	}
}

// Benchmark SecantMethod
func BenchmarkSecantMethod(b *testing.B) {
	cashFlows := createCashFlows()
	method := NewSecantMethod(1000, 1e-7)
	for i := 0; i < b.N; i++ {
		CalculateIRR(cashFlows, method)
	}
}
