package hrk_test

import (
	"fmt"
	"testing"
)

func TestNonDivisibleSubset(t *testing.T) {
	tests := []struct {
		k    int32
		s    []int32
		want int32
	}{
		//{3, []int32{1, 7, 2, 4}, 3}, // case 0
		{7, []int32{278, 576, 496, 727, 410, 124, 338, 149, 209, 702, 282, 718, 771, 575, 436}, 11}, // case 1
		/*
			The sums of all permutations of two elements from case 0 are:
			1 + 7 = 8 -> {1, 7}
			1 + 2 = 3 -> !2
			1 + 4 = 5 -> {1, 4, 7}
			7 + 2 = 9 -> !2
			7 + 4 = 11
			2 + 4 = 6 -> !2
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
	r := make(map[int32]bool)

	for i, v0 := range s {
		for n := i + 1; n < len(s); n++ {
			if vn := s[n]; (v0+vn)%3 != 0 {
				if len(r) == 0 {
					r[v0] = true
					r[vn] = true
					continue
				}
				for rk, _ := range r {
					fmt.Println("rk:", rk, "v0:", v0, "vn:", vn)
					if (v0+rk)%3 != 0 {
						fmt.Println("Add v0")
						r[v0] = true
					}
					if (vn+rk)%3 != 0 {
						fmt.Println("Add vn")
						r[vn] = true
					}
				}
			}
		}
	}

	fmt.Println(r)
	return int32(len(r))
}
