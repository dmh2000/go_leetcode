package lcjune

func reverseString(s []byte) {
	var i int
	var j int

	if len(s) <= 1 {
		return
	}

	i = 0
	j = len(s) - 1

	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
