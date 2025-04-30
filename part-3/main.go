package main

import (
	"fmt"
	"log"

	"github.com/padiazg/irr-calculator/financial"
	"github.com/padiazg/irr-calculator/performance"
	"github.com/padiazg/irr-calculator/visualization"
)

const (
	maxIter   int     = 100
	threshold float64 = 1e-6
)

func main() {
	// Define sample cash flows
	cashFlows := []financial.CashFlow{
		{Amount: -10000, Period: 0},
		{Amount: 3000, Period: 1},
		{Amount: 4000, Period: 2},
		{Amount: 5000, Period: 3},
		{Amount: 6000, Period: 4},
	}

	// Compare different root-finding methods
	methods := []struct {
		name      string
		method    financial.RootFinder
		tolerance float64
	}{
		{
			name:      "Newton-Raphson",
			method:    financial.NewNewtonRaphsonMethod(maxIter, threshold),
			tolerance: threshold,
		},
		{
			name:      "Bisection",
			method:    financial.NewBisectionMethod(maxIter, threshold),
			tolerance: threshold,
		},
		{
			name:      "Secant",
			method:    financial.NewSecantMethod(maxIter, threshold),
			tolerance: threshold,
		},
	}

	fmt.Println("IRR Calculation Comparison:")
	for _, m := range methods {
		// Perform IRR calculation
		irr, iterations := financial.CalculateIRR(cashFlows, m.method)

		fmt.Printf("%s Method:\n", m.name)
		fmt.Printf("  IRR: %.4f%%\n", irr*100)
		fmt.Printf("  Iterations: %d\n", iterations)

		// Performance profiling
		profile := performance.AnalyzeAlgorithmPerformance(
			func() {
				financial.CalculateIRR(cashFlows, m.method)
			},
			1000,
		)
		profile.AlgorithmName = m.name
		profile.PrintPerformanceReport()
		fmt.Println()
	}

	// Generate convergence visualization
	err := visualization.IRRConvergencePlot(cashFlows)
	if err != nil {
		log.Fatalf("Visualization error: %v", err)
	}
}
