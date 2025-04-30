package visualization

import (
	"os"
	"testing"

	"github.com/padiazg/irr-calculator/financial"
)

// Helper function to generate test cash flows
func generateTestCashFlows() []financial.CashFlow {
	return []financial.CashFlow{
		{Amount: -1000, Period: 0},
		{Amount: 200, Period: 1},
		{Amount: 300, Period: 2},
		{Amount: 400, Period: 3},
		{Amount: 500, Period: 4},
	}
}

// TestIRRConvergencePlot tests the IRRConvergencePlot function
func TestIRRConvergencePlot(t *testing.T) {
	cashFlows := generateTestCashFlows()

	err := IRRConvergencePlot(cashFlows)
	if err != nil {
		t.Errorf("IRRConvergencePlot returned an error: %v", err)
	}

	// Check if the plot file is created
	if _, err := os.Stat("root_finding_convergence.png"); os.IsNotExist(err) {
		t.Errorf("Expected plot file not created")
	} else {
		// Clean up the generated file after test
		os.Remove("root_finding_convergence.png")
	}
}

// BenchmarkIRRConvergencePlot benchmarks the IRRConvergencePlot function
func BenchmarkIRRConvergencePlot(b *testing.B) {
	cashFlows := generateTestCashFlows()

	for i := 0; i < b.N; i++ {
		IRRConvergencePlot(cashFlows)
	}
}
