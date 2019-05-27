package numbermath

import (
	"testing"
)

func TestEulerTotientGenerator(t *testing.T) {
	var tests = []struct {
		n   int
		obj []int
	}{
		{10, []int{0, 1, 1, 2, 2, 4, 2, 6, 4, 6}},
	}

	for _, c := range tests {
		got := EulerTotientGenerator(c.n)
		for i, num := range got {
			if c.obj[i] != num {
				t.Errorf("EulerTotientGenerator(%d)=%v (want %v) ",
					c.n, c.obj, got)
			}
		}
	}
}
