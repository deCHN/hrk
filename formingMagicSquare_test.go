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
			{4, 5, 8}, //{4, 9 ,2},
			{2, 4, 1}, //{3, 5, 7},
			{1, 9, 7}, //{8, 1, 6}, -> |5 - 6| = 1
		}, 14}, // 0 + 0 + 1 = 1
	}

	for _, v := range tests {
		if get := formingMagicSquare(v.s); get != v.want {
			t.Errorf("Given %v, want %v, but %v.", v, v.want, get)
		}
	}
}

func formingMagicSquare(s [][]int32) int32 {
	sum := int32(0)
	for _, v := range s {
		sum += row(v...)
	}

	return sum
}

func row(ns ...int32) int32 {
	sum := int32(0)
	for _, v := range ns {
		sum += v
	}

	return int32(math.Abs(float64(sum - 15)))
}
