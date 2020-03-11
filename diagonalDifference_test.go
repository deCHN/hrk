package hrk

import (
	"math"
	"testing"
)

func TestDiagonalDifference(t *testing.T) {
	input := [][]int32{
		[]int32{11, 2, 4},
		[]int32{4, 5, 6},
		[]int32{10, 8, -12},
	}
	const (
		expect = (11 + 5 + -12 - 4 - 5 - 10) * -1 //abs negative
	)

	if get := diagonalDifference(input); get != expect {
		t.Errorf("Expect %v, but %v.", 15, get)
	}
}

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	n := len(arr)
	var left, right int32

	for k, v := range arr {
		left += v[k]
		right += v[n-k-1]
	}
	return int32(math.Abs(float64(left - right)))
}
