package hrk_test

import (
	"sort"
	"testing"
)

func TestPickingNumbers(t *testing.T) {
	tests := []struct {
		a    []int32
		want int32
	}{
		{[]int32{4, 6, 5, 3, 3, 1}, 3},
		{[]int32{1, 2, 2, 3, 1, 2}, 5},
		{[]int32{6, 6, 6, 6, 6, 6}, 6}, // testcase 6
	}

	for _, v := range tests {
		if get := pickingNumbers(v.a); get != v.want {
			t.Errorf("Given: %v, want: %v, but: %v.\n", v.a, v.want, get)
		}
	}
}

func pickingNumbers(a []int32) int32 {
	max := 0

	//{4, 6, 5, 3, 3, 1},
	sort.Slice(a, func(i int, j int) bool { return a[i] < a[j] })
	//{1, 3, 3, 4, 5, 6},

	diff := func(s []int32, v0 int32) int {
		for k, v := range s {
			s[k] = v - v0
		} //v0 = 3 -> {-2, 0, 0, 1, 2, 3},

		il, ir, left := len(s), 0, int32(0) //index left / right, left(min) value of the set
		for k, v := range s {
			if v >= -1 {
				il = k
				left = v // left = 0
				break
			}
		} // il = 1

		for k, v := range s {
			if v < left+2 { // left+2 <- right(max) value of the set
				continue
			}
			ir = k - 1
			break
		} // ir = 3

		if d := ir - il; d >= 0 {
			return d + 1 // 3
		}

		return 1
	}

	size := len(a)
	for _, v0 := range a {
		s := make([]int32, size)
		copy(s, a)
		if d := diff(s, v0); max < d {
			max = d
		}
	}

	return int32(max)
}
