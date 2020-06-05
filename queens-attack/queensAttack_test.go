package hrk_test

import "testing"

func TestQueensAttack(t *testing.T) {
	tests := []struct {
		n, k      int32
		r_q, c_q  int32
		obstacles [][]int32
		want      int32
	}{
		{4, 0, 4, 4, [][]int32{nil, nil}, 9},
	}

	for _, v := range tests {
		if get := queensAttack(v.n, v.k, v.r_q, v.c_q, v.obstacles); get != v.want {
			t.Errorf("Given %v, want %v, but get %v.\n", v.n, v.want, get)
		}
	}
}

func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	return 0
}
