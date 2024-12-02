package main

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

// NewtonRaphson struct holds the parameters for the Newton-Raphson method
type NewtonRaphson struct {
	threshold    float64
	maxiter      uint
	showDetail   bool
	summaryTable table.Writer
}

// Calc performs the Newton-Raphson method to find the root of the function
func (s NewtonRaphson) Calc(value float64) float64 {
	var (
		f              = func(x float64) float64 { return (x * x) - value }
		d              = func(x float64) float64 { return 2 * x }
		x0     float64 = value
		x1     float64
		i      uint
		t      table.Writer
		result float64 = -1.0
	)

	if s.showDetail {
		fmt.Printf("Newton-Raphson | m=%.7f | threshold: %.7f | max-iter: %d\n", value, s.threshold, s.maxiter)
		t = table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"i", "xn", "xn+1", "diff"})
		t.AppendSeparator()
	}

	for ; i < s.maxiter; i++ {

		// Apply formulae
		x1 = x0 - (f(x0) / d(x0))

		// Calculate error
		diff := x0 - x1

		if t != nil {
			t.AppendRow([]interface{}{
				i,
				text.AlignRight.Apply(fmt.Sprintf("%.8f", x0), 24),
				text.AlignRight.Apply(fmt.Sprintf("%.8f", x1), 24),
				text.AlignRight.Apply(fmt.Sprintf("%.8f", diff), 24),
			})
		}

		// Check if the root is found
		if diff <= s.threshold {
			result = x1
			break
		}

		// Set the next iteration value
		x0 = x1
	}

	// Append to summary table if available
	if s.summaryTable != nil {
		s.summaryTable.AppendRow([]interface{}{
			value,
			"Newton-Raphson",
			text.AlignRight.Apply(fmt.Sprintf("%.7f", result), 24),
			i,
		})
	}

	// Render details table if available
	if t != nil {
		t.Render()
	}

	return result
}

// NewNewtonRaphson creates a new instance of the Newton-Raphson method
func NewNewtonRaphson(threshold float64, maxiter uint, showDetail bool, summaryTable table.Writer) *NewtonRaphson {
	return &NewtonRaphson{
		threshold:    threshold,
		maxiter:      maxiter,
		showDetail:   showDetail,
		summaryTable: summaryTable,
	}
}
