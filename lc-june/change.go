package lcjune

type coinSlice []int

func (a coinSlice) Len() int               { return len(a) }
func (a coinSlice) Less(i int, j int) bool { return a[j] < a[i] }
func (a coinSlice) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}

var xcount int

func changeKey(sum int, index int) int64 {
	return (int64(index) << 32) + int64(sum)
}

func changeR(amount int, coins []int, sum int, index int, m map[int64]int) int {
	xcount++

	s := (int64(index) << 32) + int64(sum)
	v, ok := m[s]
	if ok {
		return v
	}

	count := 0
	for sum < amount {
		sum += coins[index]

		if sum > amount {
			break
		}

		if sum == amount {
			count++
			break
		}

		for i := index - 1; i >= 0; i-- {
			if sum+coins[i] <= amount {
				count += changeR(amount, coins, sum, i, m)
			}
		}
	}

	m[s] = count

	return count
}

func change1(amount int, coins []int) int {
	var count int
	var m map[int64]int

	if amount == 0 {
		return 1
	}

	m = make(map[int64]int)

	xcount = 0
	count = 0
	for i := len(coins) - 1; i >= 0; i-- {
		if coins[i] == amount {
			count++
		} else if coins[i] < amount {
			count += changeR(amount, coins, 0, i, m)
		}
	}

	// for key, value := range m {
	// 	fmt.Printf("%08x,%v\n", key, value)
	// }
	return count
}

func change2(amount int, coins []int) int {
	var a []int
	a = make([]int, amount+1)

	a[0] = 1

	k := 0
	for _, coin := range coins {
		for i := coin; i < amount+1; i++ {
			k++
			a[i] += a[i-coin]
		}
	}
	return a[amount]
}
