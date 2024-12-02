package main

import (
	"fmt"
	"math"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

// Secant struct holds the parameters for the Secant method
type Secant struct {
	threshold    float64
	maxiter      uint
	showDetail   bool
	summaryTable table.Writer
}

// Calc performs the Secant method to find the root of the function
func (s Secant) Calc(value float64) float64 {
	var (
		f              = func(x float64) float64 { return (x * x) - value }
		x0     float64 = value
		x1     float64 = value - 1 // Initial guess slightly different from x0
		x2     float64
		i      uint
		t      table.Writer
		result float64 = -1.0
	)

	if s.showDetail {
		fmt.Printf("Secant | m=%.7f | x0=%.7f | x1=%.7f | threshold: %.7f | max-iter: %d\n", value, x0, x1, s.threshold, s.maxiter)
		t = table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"i", "x0", "x1", "x2", "diff"})
		t.AppendSeparator()
	}

	for ; i < s.maxiter; i++ {
		// Apply formulae
		x2 = x1 - f(x1)*(x1-x0)/(f(x1)-f(x0))

		// Calculate error
		diff := math.Abs(x2 - x1)

		if t != nil {
			t.AppendRow([]interface{}{
				i,
				text.AlignRight.Apply(fmt.Sprintf("%.8f", x0), 24),
				text.AlignRight.Apply(fmt.Sprintf("%.8f", x1), 24),
				text.AlignRight.Apply(fmt.Sprintf("%.8f", x2), 24),
				text.AlignRight.Apply(fmt.Sprintf("%.8f", diff), 24),
			})
		}

		// Check if the root is found
		if diff < s.threshold {
			result = x2
			break
		}

		// Set the next iteration values
		x0 = x1
		x1 = x2
	}

	// Append to summary table if available
	if s.summaryTable != nil {
		s.summaryTable.AppendRow([]interface{}{
			value,
			"Secant",
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

// NewSecant creates a new instance of the Secant method
func NewSecant(threshold float64, maxiter uint, showDetail bool, summaryTable table.Writer) *Secant {
	return &Secant{
		threshold:    threshold,
		maxiter:      maxiter,
		showDetail:   showDetail,
		summaryTable: summaryTable,
	}
}
