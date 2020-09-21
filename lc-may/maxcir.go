package lcmay

import "fmt"

func printCirc(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%3d,", a[i])
	}
	fmt.Println()
}

func maxSubArray(nums []int) int {
	var sum int
	var big int
	sum = nums[0]
	big = sum
	for i := 1; i < len(nums); i++ {
		n := nums[i]
		m := sum + n
		if n > 0 {
			// sum will go up
			if sum < 0 {
				// discard sum and start over
				sum = n
			} else {
				// increased, keep going
				sum = m
			}
		} else {
			// sum will go down
			if m < 0 {
				// discard and start over
				sum = n
			} else {
				// still positive, keep going
				sum = m
			}
		}

		if sum > big {
			big = sum
		}
	}
	return big
}

func maxSubarraySumCircular1(a []int) int {
	var sum1 int
	var sum2 int
	var b []int

	// first check for all negative and find largest
	maxn := -2147483648
	for i := 0; i < len(a); i++ {
		if a[i] > 0 {
			maxn = 1
			break
		}
		if a[i] <= 0 && a[i] > maxn {
			maxn = a[i]
		}
	}
	if maxn < 0 {
		return maxn
	}

	sum1 = maxSubArray(a)

	sum2 = 0
	b = make([]int, len(a))
	copy(b, a)
	for i := 0; i < len(b); i++ {
		sum2 += b[i]
		b[i] = -b[i]
	}

	sum2 = sum2 + maxSubArray(b)

	if sum2 > sum1 {
		sum1 = sum2
	}

	return sum1
}

func maxSubarraySumCircular2(a []int) int {
	var min1 int
	var min2 int
	var max1 int
	var max2 int
	var total int

	min1 = 2147483647
	min2 = 2147483647
	max1 = -2147483648
	max2 = -2147483648
	total = 0
	for i := 0; i < len(a); i++ {

		// min subarray sum
		min1 += a[i]
		if a[i] < min1 {
			min1 = a[i]
		}
		if min1 < min2 {
			min2 = min1
		}

		// max subarray sum
		max1 += a[i]
		if a[i] > max1 {
			max1 = a[i]
		}
		if max1 > max2 {
			max2 = max1
		}

		total += a[i]
	}

	if total != min2 {
		// if total not equal to min2, then
		// use max of max subarray sum or total sum - min subarray sum
		if max2 > (total - min2) {
			total = max2
		} else {
			total -= min2
		}
	} else {
		// if both are equal return the max subarray sum
		// indicates all elements are negative
		total = max2
	}

	return total
}

func maxSubarraySumCircular(a []int) int {
	var min1 int
	var min2 int
	var max1 int
	var max2 int
	var total int

	// first check for all negative and find largest
	maxn := -2147483648
	for i := 0; i < len(a); i++ {
		if a[i] > 0 {
			maxn = 1
			break
		}
		if a[i] <= 0 && a[i] > maxn {
			maxn = a[i]
		}
	}
	if maxn < 0 {
		return maxn
	}

	// compute min subarray sum and max subarray sum
	min1 = 2147483647
	min2 = 2147483647
	max1 = -2147483648
	max2 = -2147483648
	total = 0

	for i := 0; i < len(a); i++ {
		// min subarray sum
		min1 += a[i]
		if a[i] < min1 {
			min1 = a[i]
		}
		if min1 < min2 {
			min2 = min1
		}
	}

	for i := 0; i < len(a); i++ {
		// max subarray sum
		max1 += a[i]
		if a[i] > max1 {
			max1 = a[i]
		}
		if max1 > max2 {
			max2 = max1
		}

		total += a[i]
	}

	// result is greater of total - min subarray sum
	// or max sumarray sum
	total = total - min2
	if max2 > total {
		total = max2
	}

	return total
}
