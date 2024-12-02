package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

// Bisection struct holds the parameters for the Bisection method
type Bisection struct {
	threshold    float64
	maxiter      uint
	showDetail   bool
	summaryTable table.Writer
}

// Calc performs the Bisection method to find the root of the function
func (s Bisection) Calc(value float64) float64 {
	var (
		a      = 0.0
		b      = value
		sqrt   = func(x float64) float64 { return math.Pow(x, 2) - value }
		i      uint
		t      table.Writer
		result float64 = -1.0
		eval   string
	)

	if sqrt(a)*sqrt(b) >= 0 {
		log.Printf("Bisection will not work with %f:%f", a, b)
		return -1.0
	}

	if s.showDetail {
		fmt.Printf("Bisection | m=%.7f | a=%.7f | b=%.7f | threshold: %.7f | max-iter: %d\n", value, a, b, s.threshold, s.maxiter)
		t = table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"i", "a", "r", "b", "f(a)", "f(r)", "f(r)*f(a)", "eval"})
		t.AppendSeparator()
	}

	for ; i < s.maxiter; i++ {
		var (
			r   = (a + b) / 2
			fr  = sqrt(r)
			fa  = sqrt(a)
			fra = fr * fa
		)

		// Adjust bracket values
		if fra < 0 {
			b = r
			eval = "r => b"
		} else {
			a = r
			eval = "a <= r"
		}

		if t != nil {
			t.AppendRow([]interface{}{
				i,
				text.AlignRight.Apply(fmt.Sprintf("%.8f", a), 24),
				text.AlignRight.Apply(fmt.Sprintf("%.8f", r), 24),
				text.AlignRight.Apply(fmt.Sprintf("%.8f", b), 24),
				text.AlignRight.Apply(fmt.Sprintf("%.8f", fa), 24),
				text.AlignRight.Apply(fmt.Sprintf("%.8f", fr), 32),
				text.AlignRight.Apply(fmt.Sprintf("%.8f", fra), 42),
				eval,
			})
		}

		// Check if the root is found
		if fr == 0.0 || (b-a) <= s.threshold {
			result = r
			break
		}
	}

	// Append to summary table if available
	if s.summaryTable != nil {
		s.summaryTable.AppendRow([]interface{}{
			value,
			"Bisection",
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

// NewBisection creates a new instance of the Bisection method
func NewBisection(threshold float64, maxiter uint, showDetail bool, summaryTable table.Writer) *Bisection {
	return &Bisection{
		threshold:    threshold,
		maxiter:      maxiter,
		showDetail:   showDetail,
		summaryTable: summaryTable,
	}
}
