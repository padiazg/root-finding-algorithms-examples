package financial

// CashFlow represents a single cash flow with its period
type CashFlow struct {
	Amount float64
	Period int
}

// RootFinder interface for different root-finding methods
type RootFinder interface {
	FindRoot(func(float64) float64, float64) (float64, int)
}
