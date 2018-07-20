package BestTimeToBuyAndSellStock_121

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		profit := prices[i] - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}

		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}

	return maxProfit
}