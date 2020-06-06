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
		{5, 3, 4, 3, [][]int32{[]int32{5, 5}, []int32{4, 2}, []int32{2, 3}}, 10},
	}

	for _, v := range tests {
		if get := queensAttack(v.n, v.k, v.r_q, v.c_q, v.obstacles); get != v.want {
			t.Errorf("Given %v, want %v, but get %v.\n", v.n, v.want, get)
		}
	}
}

type point [2]int32

func (p point) move(d dir) point {
	return d.next(p)
}

type dir interface {
	next(p point) point
}

func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	orgin := point{r_q, c_q}

	moves := 0

	for d := range directions {

	}

	return 0
}
