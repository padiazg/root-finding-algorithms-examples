package visualization

import (
	"log"
	"math"

	"github.com/padiazg/irr-calculator/financial"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// IRRConvergencePlot generates a visualization of algorithm convergence
func IRRConvergencePlot(cashFlows []financial.CashFlow) error {
	p := plot.New()

	p.Title.Text = "Root Finding Methods Convergence"
	p.X.Label.Text = "Iteration"
	p.Y.Label.Text = "IRR Value"

	// Define root-finding methods
	methods := []struct {
		name   string
		method financial.RootFinder
	}{
		{
			name:   "Newton-Raphson",
			method: financial.NewNewtonRaphsonMethod(20, 1e-6), //&financial.NewtonRaphsonMethod{MaxIterations: 20, Tolerance: 1e-6},
		},
		{
			name:   "Bisection",
			method: financial.NewBisectionMethod(20, 1e-6), //&financial.BisectionMethod{MaxIterations: 20, Tolerance: 1e-6},
		},
		{
			name:   "Secant",
			method: financial.NewSecantMethod(20, 1e-6),
		},
	}

	// Convergence tracking plot points
	var plotPoints []plotter.XYs

	for _, m := range methods {
		methodPoints := make(plotter.XYs, 0)

		npvFunc := func(rate float64) float64 {
			npv := 0.0
			for _, cf := range cashFlows {
				npv += cf.Amount / math.Pow(1+rate, float64(cf.Period))
			}
			return npv
		}

		currentRate := 0.1
		tolerance := 1e-6
		maxIterations := 20

		switch m.method.(type) {
		case *financial.NewtonRaphsonMethod:
			for i := 0; i < maxIterations; i++ {
				fx := npvFunc(currentRate)
				if math.Abs(fx) < tolerance {
					break
				}

				h := 1e-7
				dfx := (npvFunc(currentRate+h) - npvFunc(currentRate)) / h

				if dfx == 0 {
					break
				}

				currentRate = currentRate - fx/dfx
				methodPoints = append(methodPoints, plotter.XY{X: float64(i), Y: currentRate})
			}

		case *financial.BisectionMethod:
			lower, upper := 0.0, 1.0
			for i := 0; i < maxIterations; i++ {
				midRate := (lower + upper) / 2
				fmid := npvFunc(midRate)

				if math.Abs(fmid) < tolerance {
					break
				}

				flow := npvFunc(lower)
				if flow*fmid < 0 {
					upper = midRate
				} else {
					lower = midRate
				}

				methodPoints = append(methodPoints, plotter.XY{X: float64(i), Y: midRate})
			}

		case *financial.SecantMethod:
			x0, x1 := 0.1, 0.2
			for i := 0; i < maxIterations; i++ {
				fx0 := npvFunc(x0)
				fx1 := npvFunc(x1)

				if math.Abs(fx1) < tolerance {
					break
				}

				x2 := x1 - fx1*(x1-x0)/(fx1-fx0)
				methodPoints = append(methodPoints, plotter.XY{X: float64(i), Y: x2})

				x0, x1 = x1, x2
			}
		}

		plotPoints = append(plotPoints, methodPoints)
	}

	// Create plot with points
	plotArgs := []interface{}{"Newton-Raphson", plotPoints[0],
		"Bisection", plotPoints[1],
		"Secant", plotPoints[2]}

	err := plotutil.AddLinePoints(p, plotArgs...)
	if err != nil {
		return err
	}

	// Save plot
	if err := p.Save(8*vg.Inch, 6*vg.Inch, "root_finding_convergence.png"); err != nil {
		log.Fatal(err)
	}

	return nil
}
