package hrk_test

import (
	"fmt"
	"sort"
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
		if get := nonDivisibleSubset2(v.k, v.s); get != v.want {
			t.Errorf("Given %v, want %v, but get %v.\n", v.s, v.want, get)
		}
	}
}

// https://www.hackerrank.com/challenges/non-divisible-subset/problem
// Given a set of distinct integers, print the size of a maximal subset of S where
// the sum of any 2 numbers in S' is not evenly divisible by k.
func nonDivisibleSubset(k int32, s []int32) int32 {
	mx := int32(0)

	for i, _ := range s {
		si := make([]int32, len(s))
		copy(si, s)
		si[0], si[i] = si[i], si[0]
		fmt.Println(si)
		if x := subset(k, si); x > mx {
			mx = x
		}
	}

	return mx
}

func printmap(m map[int32]bool) {
	keys := []int32{}
	for k, _ := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i int, j int) bool { return keys[i] < keys[j] })

	fmt.Println("\t", keys, "len:", len(keys))
}

func subset(k int32, s []int32) int32 {
	r := make(map[int32]bool)

	for i, v0 := range s {
		for n := i + 1; n < len(s); n++ {
			if vn := s[n]; (v0+vn)%3 != 0 {
				if len(r) == 0 {
					r[v0] = true
					r[vn] = true
					continue
				}
				ok0, okn := false, false
				for rk, _ := range r {
					if (v0+rk)%3 == 0 {
						ok0 = true
						break
					}
				}

				for rn, _ := range r {
					if (vn+rn)%3 == 0 {
						okn = true
						break
					}
				}

				if !ok0 {
					r[v0] = true
				}

				if !okn {
					r[vn] = true
				}
			}
		}
	}

	printmap(r)

	return int32(len(r))
}

func nonDivisibleSubset2(k int32, s []int32) int32 {
	r := make(map[int32]int32)

	for _, v := range s {
		r[v%k]++
	}

	mx := int32(0)
	for i := int32(1); i < int32(len(r)); i++ {
		if x := sum(i, r); x > mx {
			mx = x
		}
	}
	return mx + r[0]
}

func sum(i int32, r map[int32]int32) int32 {
	sum := int32(0)
	for ir, v := range r {
		if ir != i {
			sum += v
		}
	}
	return sum
}
