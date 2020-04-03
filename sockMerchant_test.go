package hrk_test

import "testing"

func TestSockMerchant(t *testing.T) {
	td := []struct {
		n      int32
		ar     []int32
		expect int32
	}{
		{7, []int32{1, 2, 1, 2, 1, 3, 2}, 2},
		{9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}, 3},
	}

	for _, v := range td {
		if get := sockMerchant(v.n, v.ar); get != v.expect {
			t.Errorf("Expect %v, but %v.", v.expect, get)
		}
	}
}

// Constraints:
// 1 < n < 100
// 1 <= ar[i] <=100 where 0 <= i <= n
func sockMerchant(n int32, ar []int32) int32 {
	pairs := [101]int32{}

	for _, v := range ar {
		if pairs[v] == 0 {
			pairs[v] = 1
		} else {
			pairs[v] = 0
		}
	}

	var sum int32

	for _, v := range pairs {
		sum += v
	}

	return (n - sum) / 2
}
