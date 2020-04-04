package hrk_test

import "testing"

func TestPageCount(t *testing.T) {
	td := []struct {
		n, p   int32
		expect int32
	}{
		{6, 2, 1},
		{5, 4, 0},
	}

	for _, v := range td {
		if get := pageCount(v.n, v.p); get != v.expect {
			t.Errorf("Expect %v, but %v.", v.expect, get)
		}
	}

}

// Constraints:
// 1 <= n <= 10^5
// 1 <= p <= n
func pageCount(n int32, p int32) int32 {

}
