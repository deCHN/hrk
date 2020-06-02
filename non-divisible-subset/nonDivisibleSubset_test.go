package hrk_test

import (
	"testing"
)

func TestNonDivisibleSubset(t *testing.T) {
	tests := []struct {
		k    int32
		s    []int32
		want int32
	}{
		//{3, []int32{1, 7, 2, 4}, 3}, // case 0
		//{7, []int32{278, 576, 496, 727, 410, 124, 338, 149, 209, 702, 282, 718, 771, 575, 436}, 11}, // case 1
		//{1, []int32{1, 2, 3, 4, 5}, 1}, // case 6
		{4, []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5}, // case 7
	}

	for _, v := range tests {
		if get := nonDivisibleSubset(v.k, v.s); get != v.want {
			t.Errorf("Given %v, want %v, but get %v.\n", v.s, v.want, get)
		}
	}
}

// https://www.hackerrank.com/challenges/non-divisible-subset/problem
// Given a set of distinct integers, print the size of a maximal subset of S where
// the sum of any 2 numbers in S' is not evenly divisible by k.
// 1 <= k <= 100
func nonDivisibleSubset(k int32, s []int32) int32 {
	if k == 1 {
		return 1
	}

	r := make(map[int32]int32)

	for _, v := range s {
		r[v%k]++
	}

	set := int32(0)
	skip := make(map[int32]bool)

	for i, vi := range r {
		if _, ok := skip[i]; ok {
			continue
		}

		skip[i] = true

		if vri, ok := r[k-i]; ok {
			if vi > vri {
				set += vi
			} else {
				set += vri
			}
			skip[k-i] = true
		} else {
			set += vi
		}
	}

	return set
}
