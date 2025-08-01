package symdel

func DemerauLevenshteinDistance(x, y Entry) float64 {
	a, b := []rune(x.Word), []rune(y.Word)
	if len(a) > len(b) {
		a, b = b, a
	}
	start, la, lb := 0, len(a), len(b)
	for la > 0 && a[la-1] == b[lb-1] {
		la--
		lb--
	}
	for start < la && a[start] == b[start] {
		start++
	}
	if la == start {
		return float64(lb - start)
	}
	a, b = a[start:la], b[start:lb]
	la -= start
	lb -= start
	d := make([]int, lb)
	for i := 0; i < lb; i++ {
		d[i] = i + 1
	}
	current := 0
	for i := 0; i < la; i++ {
		c := a[i]
		left := i
		current = i
		for j := 0; j < lb; j++ {
			above := current
			current = left
			left = d[j]
			if c != b[j] {
				current = min(min(above, current), left) + 1
			}
			d[j] = current
		}
	}
	return float64(current)
}
