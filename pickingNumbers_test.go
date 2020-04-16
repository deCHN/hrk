package hrk_test

import "testing"

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
	return 0
}
