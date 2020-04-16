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

	//size := len(a)
	//n := 0

	diff := func(s []int32, v0 int32) int {
		for k, v := range s {
			s[k] = v - v0
		}
		//v0 = 3 -> {-2, 0, 0, 1, 2, 3},
		//TODO: find first index of -1 or 0
		//TODO: find last index of 0 or 1

		return 0
	}

	// {1,2
	for _, v0 := range a {
		s := make([]int32, 0)
		copy(s, a)
		diff(s, v0)
	}

	return int32(max)
}
