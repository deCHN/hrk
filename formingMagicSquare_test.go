package hrk_test

import (
	"math"
	"testing"
)

func TestFormingMagicSquare(t *testing.T) {

	tests := []struct {
		s    [][]int32
		want int32
	}{
		{[][]int32{
			{5, 3, 4}, //{8, 3, 4}, -> |5 - 8| = 3
			{1, 5, 8}, //{1, 5, 9}, -> |8 - 9| = 1
			{6, 4, 2}, //{6, 7 ,2}, -> |4 - 7| = 3
		}, 7}, // 3 + 1 + 3 = 7
		{[][]int32{
			{4, 9, 2}, //{4, 9 ,2},
			{3, 5, 7}, //{3, 5, 7},
			{8, 1, 5}, //{8, 1, 6}, -> |5 - 6| = 1
		}, 1}, // 0 + 0 + 1 = 1
		{[][]int32{
			{4, 5, 8}, //{2, 9, 4}, ->
			{2, 4, 1}, //{7, 5, 3},
			{1, 9, 7}, //{6, 1, 8},
		}, 14}, //
		{[][]int32{
			{2, 9, 8}, //{6, 7, 2},
			{4, 2, 7}, //{1, 5, 9},
			{5, 6, 7}, //{8, 3, 4},
		}, 21}, //
	}

	for _, v := range tests {
		if get := formingMagicSquare(v.s); get != v.want {
			t.Errorf("Given %v, want %v, but %v.", v, v.want, get)
		}
	}
}

func formingMagicSquare(s [][]int32) int32 {

	//cover from two dimentions to one for easy calculate
	covert := func(s [][]int32) []int32 {
		ret := make([]int32, 9)
		for _, v := range s {
			ret = append(ret, v...)
		}
		return ret
	}

	minCost := func(s []int32) int32 {
		squares := [8][9]int32{
			[9]int32{2, 9, 4, 7, 5, 3, 6, 1, 8},
			[9]int32{4, 9, 2, 3, 5, 7, 8, 1, 6}, //mirrow of last one
			[9]int32{6, 7, 2, 1, 5, 9, 8, 3, 4}, //rotate 90 degrees of first one
			[9]int32{2, 7, 6, 9, 5, 1, 4, 3, 8}, //mirrow of last one
			[9]int32{8, 1, 6, 3, 5, 7, 4, 9, 2}, //rotate 90 degrees of third one
			[9]int32{6, 1, 8, 7, 5, 3, 2, 9, 4}, //mirrow of last one
			[9]int32{4, 3, 8, 9, 5, 1, 2, 7, 6}, //rotate 90 degrees of third one
			[9]int32{8, 3, 4, 1, 5, 9, 6, 7, 2}, //mirrow of last one
		}

		cost := float64(100)
		for _, sqr := range squares {
			sum := float64(0)
			for k, v := range sqr {
				sum += math.Abs(float64(s[k] - v))
			}
			if sum < cost {
				cost = sum
			}
		}
		return int32(cost)
	}

	return minCost(covert(s))
}
