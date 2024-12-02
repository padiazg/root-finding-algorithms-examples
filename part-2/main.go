package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

const (
	threshold float64 = 0.0000001
	maxiter   uint    = 64
)

func main() {
	var (
		showDetails = flag.Bool("show-details", false, "Show detailed output")
		showSummary = flag.Bool("show-summary", false, "Show results summary")
		valuesStr   = flag.String("values", "2", "Values")
		summary     table.Writer
	)

	flag.Parse()

	if *showSummary {
		summary = table.NewWriter()
		summary.SetOutputMirror(os.Stdout)
		summary.AppendHeader(table.Row{"Value", "Algorithm", "Result", "Iterations"})
		summary.SetColumnConfigs([]table.ColumnConfig{
			{Number: 1, AutoMerge: true},
			{Number: 2, Align: text.AlignLeft, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
			{Number: 3, Align: text.AlignRight, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
			{Number: 4, Align: text.AlignRight, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
			// {Number: 5, Align: text.AlignLeft, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		})
		// summary.SetStyle(table.StyleLight)
		summary.Style().Options.SeparateRows = true

	}

	// Parse the values from the command-line argument
	values := parseValues(*valuesStr)

	for _, m := range values {
		// fmt.Printf("-------------------------\nm=%.8f\n", m)

		vb := NewBisection(threshold, maxiter, *showDetails, summary)
		vb.Calc(m)
		// fmt.Printf("bisection     : %f\n", vb.Calc(m))

		vn := NewNewtonRaphson(threshold, maxiter, *showDetails, summary)
		vn.Calc(m)
		// fmt.Printf("newton-raphson: %f\n", vn.Calc(m))

		vs := NewSecant(threshold, maxiter, *showDetails, summary)
		vs.Calc(m)
		// fmt.Printf("secant        : %f\n", vs.Calc(m))

		// fmt.Println("")
	}

	if *showSummary {
		fmt.Printf("threshold: %.7f | max-iter: %d\n", threshold, maxiter)
		summary.Render()
	}

}

func parseValues(valuesStr string) []float64 {
	strValues := strings.Split(valuesStr, ",")
	var values []float64
	for _, strValue := range strValues {
		if value, err := strconv.ParseFloat(strings.TrimSpace(strValue), 64); err == nil {
			values = append(values, value)
		}
	}
	return values
}
