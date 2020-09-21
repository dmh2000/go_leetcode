package lcmay

// StockSpanner ...
type StockSpanner struct {
	sum   int
	pidx  int
	price []int
	cidx  int
	count []int
}

// SpanConstructor ...
func SpanConstructor() StockSpanner {
	span := StockSpanner{0, 0, []int{}, 0, []int{}}
	span.price = make([]int, 15000)
	span.count = make([]int, 15000)
	return span
}

// Next ...
func (span *StockSpanner) Next(price int) int {
	span.sum++

	out := 0

	// while count array is not empty
	for span.cidx > 0 {
		prev := span.count[span.cidx] - 1
		// if this price is less than previous price, stop
		if price < span.price[prev] {
			break
		}
		// pop the count
		span.cidx--
	}

	if len(span.count) == 0 {
		out = span.sum
	} else {
		out = span.sum - span.count[span.cidx]
	}

	// push the price
	span.price[span.pidx] = price
	span.pidx++

	// push the count
	span.cidx++
	span.count[span.cidx] = span.sum

	return out
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

/*
 // StockSpanner ...
type StockSpanner struct {
	index int
	sum   int
	price []int
	count []int
}

// SpanConstructor ...
func Constructor() StockSpanner {
	span := StockSpanner{0, 0, []int{}, []int{}}
	return span
}

// Next ...
func (span *StockSpanner) Next(price int) int {
	span.sum++

	out := 0

	for len(span.count) > 0 {
		c := span.count[0]
		if price < span.price[c-1] {
			break
		}
		span.count = span.count[1:]
	}
	if len(span.count) == 0 {
		out = span.sum
	} else {
		out = span.sum - span.count[0]
	}
	span.price = append(span.price, price)
	span.count = append([]int{span.sum}, span.count...)

	return out
}
*/
