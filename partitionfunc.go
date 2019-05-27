package numbermath

import (
	m "math"
	"errors"
)

func PartitionFunc(n int) ([]int, error) {
	if n > 405 {
		msg := "numbermath: upper bound larger than 405 will" +
			   "make the result overflow."
		return []int{}, errors.New(msg)
	}
	pl := make([]int, n+1)
	pl[0] = 1
	pl[1] = 1
	for i := 2; i <= n; i++ {
		a := m.Sqrt(float64(24*n + 1))
		lowb := int(m.Floor((-a + 1) / 6))
		upb := int(m.Ceil((a + 1) / 6))
		for k := lowb; k <= upb; k++ {
			if k != 0 {
				d := i - k*(3*k-1)/2
				if d >= 0 {
					sign := 0
					if k%2 == 0 {
						sign = -1
					} else {
						sign = 1
					}
					pl[i] += (sign * pl[d])
				}
			}
		}
	}
	return pl, nil
}
