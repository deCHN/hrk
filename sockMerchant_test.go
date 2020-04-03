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

func sockMerchant(n int32, ar []int32) int32 {

}
