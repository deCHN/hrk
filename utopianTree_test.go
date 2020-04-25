package hrk_test

import "testing"

func TestUtopianTree(t *testing.T) {
	tests := []struct {
		n, want int32
	}{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 6},
		{4, 7},
		{5, 14},
	}

	for _, v := range tests {
		if get := utopianTree(v.n); get != v.want {
			t.Errorf("Given %v, want %v, but %v.\n", v.n, v.want, get)
		}
	}
}

// https://www.hackerrank.com/challenges/utopian-tree/problem
func utopianTree(n int32) int32 {
	if n == 0 {
		return 1
	} else {
		if n%2 == 0 {
			return utopianTree(n-1) + 1
		}
		return utopianTree(n-1) * 2
	}
}
