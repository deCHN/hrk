package hrk_test

import "testing"

func TestPermutationEquation(t *testing.T) {
	tests := []struct {
		p, want []int32
	}{
		{[]int32{}, []int32{}},
	}

	for _, v := range tests {
		get := permutationEquation(v.p)
		for k, item := range get {
			if item != v.want[k] {
				t.Errorf("Given %v, want %v, but get %v.\n", v.p[k], v.want[k], item)
			}
		}
	}
}

func permutationEquation(p []int32) []int32 {
	return nil
}
