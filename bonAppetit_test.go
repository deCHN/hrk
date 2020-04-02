package hrk_test

import "testing"

func TestBonAppetit(t *testing.T) {
	td := []struct {
		bill   []int32
		k      int32
		b      int32
		expect int
	}{
		{bill: []int32{3, 10, 2, 9}, k: 4, b: 12, expect: 5},
		{bill: []int32{3, 10, 2, 9}, k: 4, b: 7, expect: 0},
	}

	for _, tc := range td {
		if get := bonAppetit(tc.bill, tc.k, tc.b); get != tc.expect {
			t.Errorf("Expect %v, but %v.", tc.expect, get)
		}
	}
}

func bonAppetit(bill []int32, k int32, b int32) int {

}
