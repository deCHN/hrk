package hrk_test

import (
	"testing"
)

func TestQueensAttack(t *testing.T) {
	tests := []struct {
		n, k      int32
		r_q, c_q  int32
		obstacles [][]int32
		want      int32
	}{
		{4, 0, 4, 4, [][]int32{nil}, 9},
		{5, 3, 4, 3, [][]int32{[]int32{5, 5}, []int32{4, 2}, []int32{2, 3}}, 10},
	}

	for _, v := range tests {
		if get := queensAttack(v.n, v.k, v.r_q, v.c_q, v.obstacles); get != v.want {
			t.Errorf("Given %v, want %v, but get %v.\n", v.n, v.want, get)
		}
	}
}

// https://www.hackerrank.com/challenges/queens-attack-2/problem
func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	// row - column pair
	type point [2]int32

	type direction func(p point) point

	sw := func(p point) point { return point{p[0] - 1, p[1] - 1} }
	south := func(p point) point { return point{p[0] - 1, p[1]} }
	se := func(p point) point { return point{p[0] - 1, p[1] + 1} }
	east := func(p point) point { return point{p[0], p[1] + 1} }
	ne := func(p point) point { return point{p[0] + 1, p[1] + 1} }
	north := func(p point) point { return point{p[0] + 1, p[1]} }
	nw := func(p point) point { return point{p[0] + 1, p[1] - 1} }
	west := func(p point) point { return point{p[0], p[1] - 1} }

	// check checks if the point p is a valid location
	check := func(p point) bool {
		// within the board
		if p[0] < 1 || p[1] < 1 {
			return false
		}

		if p[0] > n || p[1] > n {
			return false
		}

		// not on obstacles
		for _, ob := range obstacles {
			if ob == nil {
				break
			}

			if p[0] == ob[0] && p[1] == ob[1] {
				return false
			}
		}

		return true
	}

	var directions []direction = []direction{sw, south, se, east, ne, north, nw, west}

	moves := int32(0)

	for _, dir := range directions {
		p := point{r_q, c_q}
		for {
			p = dir(p)
			if check(p) {
				moves++
				continue
			}
			break
		}
	}

	return moves
}
