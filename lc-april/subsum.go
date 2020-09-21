package lc

import "fmt"

func subarraySum1(nums []int, k int) int {
	counts := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			counts[sum] = counts[sum] + 1
		}
	}
	fmt.Println(counts)
	return counts[k]
}

type sumPair struct {
	i int
	j int
}

func subarraySum2(nums []int, k int) int {
	sum := make(map[int]int)
	v := 0
	count := 0
	for j := 0; j < len(nums); j++ {
		v += nums[j]
		if v == k {
			count++
		}
		sum[j] = sum[j] + v
	}
	fmt.Println(sum)

	for j := 1; j < len(nums); j++ {
		for i := 0; i < j; i++ {
			s := sum[j] - sum[i]
			fmt.Println(j, i, s)
			if s == k {
				count++
			}
		}
	}
	fmt.Println(sum)
	return count
}

func subarraySum3(nums []int, k int) int {
	sum := make(map[int]int)
	v := 0
	count := 0
	for j := 0; j < len(nums); j++ {
		v += nums[j]

		// if v matches k then count it
		if v == k {
			count++
		}

		// j is current index
		// i is v-k
		// sum[i] = sum[v-k]
		fmt.Println(sum, v-k, sum[v-k])
		count += sum[v-k]
		fmt.Printf("%v\n\n", count)

		sum[v] = sum[v] + 1
	}

	return count
}

func subarraySum(nums []int, k int) int {
	return subarraySum3(nums, k)
}
