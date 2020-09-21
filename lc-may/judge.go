package lcmay

func findJudge(N int, trust [][]int) int {

	var trusts []int
	var istrusted []int
	var a int
	var b int

	trusts = make([]int, N+1)
	istrusted = make([]int, N+1)
	for i := 0; i < len(trust); i++ {
		// a trusts someone (judge == 0)
		a = trust[i][0]
		trusts[a]++

		// b is trusted (judge == N-1)
		b = trust[i][1]
		istrusted[b]++
	}

	// find any who dont trust anyone (trusts[i] = 0)
	// find any who are trusted by all (istrusted[i] == N-1)
	for i := 1; i <= N; i++ {
		if trusts[i] == 0 && istrusted[i] == N-1 {
			return i
		}

	}
	// no judge
	return -1
}
