package lc

func maxProfit(prices []int) int {
	profit := 0
	b := 0
	s := 0
	i := 1
	for i < len(prices) {
		if prices[s] > prices[i] {
			profit += prices[s] - prices[b]
			b = i
			s = i
		} else {
			s = i
		}
		i++
	}
	if s > b {
		profit += prices[s] - prices[b]
	}
	return profit
}
