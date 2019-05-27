package numbermath

import (
	"testing"
)

func TestPartitionFunc(t *testing.T) {
	var tests = []struct {
		n   int
		obj []int
	}{
		{39, []int{1, 1, 2, 3, 5, 7, 11, 15, 22, 30, // 0~9
			42, 56, 77, 101, 135, 176, 231, 297, 385, 490, // 10~19
			627, 792, 1002, 1255, 1575, // 20~24
			1958, 2436, 3010, 3718, 4565, // 25~29
			5604, 6842, 8349, 10143, 12310, // 30~34
			14883, 17977, 21637, 26015, 31185}}, // 35~39
		{406, []int{}},
	}

	for _, c := range tests {
		got, err := PartitionFunc(c.n)
		if c.n >= 406 {
			if err == nil {
				t.Errorf("Upper bound larger than 405 should not be allowed")
			}
		} else {
			for i, num := range got {
				if c.obj[i] != num {
					t.Errorf("PartitionFunc(%d)=%v (want %v) ",
						c.n, c.obj, got)
				}
			}
		}
	}
}
