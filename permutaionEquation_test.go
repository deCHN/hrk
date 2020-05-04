package hrk_test

import (
	"testing"
)

func TestPermutationEquation(t *testing.T) {
	tests := []struct {
		p, want []int32
	}{
		{[]int32{2, 3, 1}, []int32{2, 3, 1}}, // 1 === p(3) === p(p(2)); 2 === p(1) === p(p(3)); 3 === p(2) === p(p(1))
		{[]int32{4, 3, 5, 1, 2}, []int32{1, 3, 5, 4, 2}},
	}

	for _, v := range tests {
		get := permutationEquation(v.p)
		for k, item := range v.want {
			if len(get) <= k || get[k] != item {
				t.Errorf("Given %v, want %v, but get %v.\n", v.p, v.want, get)
				break
			}
		}
	}
}

// https://www.hackerrank.com/challenges/permutation-equation/problem
func permutationEquation(p []int32) []int32 {
	var e permutation = p

	var r []int32

	for x := int32(1); x <= int32(len(e)); x++ {
		r = append(r, e.p(e.p(x)))
	}

	return r
}

type permutation []int32

func (e permutation) p(i int32) int32 {
	for k, v := range e {
		if v == i {
			return int32(k + 1)
		}
	}

	return -1
}
