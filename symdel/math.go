package symdel

func binomial(n, k int) int {
	res := 1
	m := n + 1
	terms := min(k, n-k)

	for i := 1; i <= terms; i++ {
		res *= m - i
		res /= i
	}
	return res
}

func binomialsum(start, end int) int {
	res := 0
	for i := start; i < end; i++ {
		res += binomial(end, i)
	}
	return res
}
