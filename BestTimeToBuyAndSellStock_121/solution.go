package BestTimeToBuyAndSellStock_121

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	maxProfit := 0
	maxSell := prices[len(prices)-1]
	for i := len(prices)-2; i >= 0; i-- {
		buy := prices[i]
		profit := maxSell - buy
		if profit > maxProfit {
			maxProfit = profit
		}

		if buy > maxSell {
			maxSell = buy
		}
	}

	return maxProfit
}