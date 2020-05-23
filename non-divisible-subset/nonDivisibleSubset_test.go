package hrk_test

import "testing"

func TestNonDivisibleSubset(t *testing.T) {
	tests := []struct {
		k    int32
		s    []int32
		want int32
	}{
		{3, []int32{1, 7, 2, 4}, 3}, // case 0
		/*
			The sums of all permutations of two elements from case 0 are:
			1 + 7 = 8
			1 + 2 = 3
			1 + 4 = 5
			7 + 2 = 9
			7 + 4 = 11
			2 + 4 = 6
			We see that only S'={1,7,4} will not ever sum to a multiple of k=3.
		*/
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
func nonDivisibleSubset(k int32, s []int32) int32 {
	return 0
}
