package hrk_test

import "testing"

func TestSquares(t *testing.T) {
	tests := []struct {
		a, b, want int32
	}{
		{3, 9, 2},
		{17, 24, 0},
	}

	for _, v := range tests {
		if get := squares(v.a, v.b); get != v.want {
			t.Errorf("Given %v, want %v, but get %v.\n", v, v.want, get)
		}
	}

}

// https://www.hackerrank.com/challenges/sherlock-and-squares/problem
// 1 ≤ A ≤ B ≤ 10^9
func squares(a int32, b int32) int32 {
	return 0
}
