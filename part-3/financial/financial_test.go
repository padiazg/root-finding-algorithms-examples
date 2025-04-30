package financial

// Helper function to create cash flows
func createCashFlows() []CashFlow {
	return []CashFlow{
		{Amount: -1000, Period: 0},
		{Amount: 300, Period: 1},
		{Amount: 420, Period: 2},
		{Amount: 680, Period: 3},
	}
}

// var cashFlows = []CashFlow{
// 	{Amount: -1000, Period: 0},
// 	{Amount: 300, Period: 1},
// 	{Amount: 420, Period: 2},
// 	{Amount: 680, Period: 3},
// }
