package hrk_test

import "testing"

func TestPermutationEquation(t *testing.T) {
	tests := []struct {
		p, want []int32
	}{
		{[]int32{}, []int32{}},
	}

	for _, v := range tests {
		if get := permutationEquation(v.p); get != v.want {
			t.Errorf("Given %v, want %v, but get %v.\n", v, v.want, get)
		}
	}
}

func permutationEquation(p []int32) []int32 {
	return []int32{}
}
