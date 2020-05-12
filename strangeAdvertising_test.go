package hrk_test

import (
	"math"
	"testing"
)

func TestViralAdvertising(t *testing.T) {
	tests := []struct {
		n, want int32
	}{
		{1, 2},
		{2, 5},
		{3, 9},
		{4, 15},
		{5, 24},
	}

	for _, v := range tests {
		if get := viralAdvertising(v.n); get != v.want {
			t.Errorf("Given %v, want %v, but %v.\n", v.n, v.want, get)
		}
	}
}

/*
* https://www.hackerrank.com/challenges/strange-advertising/problem
*
* Day Shared Liked Cumulative
* 1      5     2       2
* 2      6     3       5
* 3      9     4       9
* 4     12     6      15
* 5     18     9      24
 */
func viralAdvertising(n int32) int32 {

	sum := int32(0)

	for i := int32(1); i <= n; i++ {
		sum += liked(i)
	}

	return sum
}

func liked(n int32) int32 {
	if n == 1 {
		return 2
	}

	return int32(math.Floor(float64(liked(n-1)) * 3 / 2))
}
