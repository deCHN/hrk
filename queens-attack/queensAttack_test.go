package hrk_test

import (
	"fmt"
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
		//{5, 3, 4, 3, [][]int32{[]int32{5, 5}, []int32{4, 2}, []int32{2, 3}}, 10},
	}

	for _, v := range tests {
		if get := queensAttack(v.n, v.k, v.r_q, v.c_q, v.obstacles); get != v.want {
			t.Errorf("Given %v, want %v, but get %v.\n", v.n, v.want, get)
		}
	}
}

// row - column pair
type point [2]int32

type direction func(p point) point

func SW(p point) point { return point{p[0] - 1, p[1] - 1} }
func S(p point) point  { return point{p[0] - 1, p[1]} }
func SE(p point) point { return point{p[0] - 1, p[1] + 1} }
func E(p point) point  { return point{p[0], p[1] + 1} }
func NE(p point) point { return point{p[0] + 1, p[1] + 1} }
func N(p point) point  { return point{p[0] + 1, p[1]} }
func NW(p point) point { return point{p[0] + 1, p[1] - 1} }
func W(p point) point  { return point{p[0], p[1] - 1} }

func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	//check checks if the point p is a valid location.
	check := func(p point) bool {
		if p[0] > n || p[1] > n {
			return false
		}

		for _, ob := range obstacles {
			if ob == nil {
				fmt.Println("ob break")
				break
			}

			if p[0] == ob[0] && p[1] == ob[1] {
				return false
			}
		}

		return true
	}

	var directions []direction = []direction{SW, S, SE, E, NE, N, NW, W}

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
