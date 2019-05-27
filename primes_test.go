package numbermath

import (
	"testing"
)

func TestPrimeGenerator(t *testing.T) {
	var tests = []struct {
		n   int
		obj []int
	}{
		{12, []int{2, 3, 5, 7, 11}},
		{14, []int{2, 3, 5, 7, 11, 13}},
		{12, []int{2, 3, 5, 7, 11}},
	}

	for _, c := range tests {
		i := 0
		for num := range PrimeGenerator(c.n) {
			if c.obj[i] != num {
				t.Errorf("PrimeGenerator error: %d", c.n)
			}
			i++
		}
	}
}
