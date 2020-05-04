package hrk_test

import "testing"

func TestPermutationEquation(t *testing.T) {
	tests := []struct {
		p, want []int32
	}{
		{[]int32{2, 3, 1}, []int32{2, 3, 1}},
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

func permutationEquation(p []int32) []int32 {
	return []int32{}
}
