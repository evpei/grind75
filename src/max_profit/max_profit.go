package maxprofit

import "math"

// You are given an array prices where prices[i] is the price of a given stock on the ith day.
//
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
//
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

func Solve(prices []int) int {
	// iterate over the prices, return biggest combination price[n] + price[n + x] where x > 0, otherwise return 0
	maxProfit := 0
	minBuyPrice := math.MaxInt

	for i := 0; i < len(prices)-1; i++ {
		if minBuyPrice > prices[i] {
			minBuyPrice = prices[i]
		}

		currentProfit := prices[i+1] - minBuyPrice

		if maxProfit < currentProfit {
			maxProfit = currentProfit
		}
	}

	return maxProfit
}
