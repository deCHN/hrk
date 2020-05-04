package hrk_test

import "testing"

func TestPermutationEquation(t *testing.T) {
	tests := []struct {
		p, want []int32
	}{
		{[]int32{}, []int32{1}},
	}

	for _, v := range tests {
		get := permutationEquation(v.p)
		for k, item := range v.want {
			if get[k] != item {
				t.Errorf("Given %v, want %v, but get %v.\n", v.p[k], item, get[k])
			}
		}
	}
}

func permutationEquation(p []int32) []int32 {
	return nil
}
