package hrk_test

import (
	"math"
	"testing"
)

func TestPageCount(t *testing.T) {
	td := []struct {
		n, p   int32
		expect int32
	}{
		{6, 2, 1},
		{5, 4, 0},
		{1, 1, 0},
		{7, 7, 0},
		{100, 100, 0},
	}

	for _, v := range td {
		if get := pageCount(v.n, v.p); get != v.expect {
			t.Errorf("Given: %v, expect %v, but %v.", v, v.expect, get)
		}
	}

}

// Constraints:
// 1 <= n <= 10^5
// 1 <= p <= n
func pageCount(n int32, p int32) int32 {
	//turn[0] = 1
	//turn[1] = 2, 3
	//turn[2] = 4, 5
	//turn[3] = 6, 7
	//turn[k] = 2k, 2k+1
	//turn...
	//turn[total] = (n + 1) / 2

	k := math.Floor(float64(p) / 2)
	total := math.Ceil(float64(n+1) / 2)
	if back := total - k - 1; back < k {
		return int32(back)
	}
	return int32(k)
}
