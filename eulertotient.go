package numbermath

func EulerTotientGenerator(under int) []int {
	sieve := make([]int, under)
	for i := range sieve {
		if i >= 2 {
			sieve[i] = i - 1
		} else {
			sieve[i] = i
		}
	}
	for i, _ := range sieve {
		if i >= 2 {
			for b := 2 * i; b < under; b += i {
				sieve[b] -= sieve[i]
			}
		}
	}
	return sieve
}
