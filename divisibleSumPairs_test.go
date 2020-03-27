package hrk

import "testing"

type testdata struct {
	n, k   int32
	ar     []int32
	expect int32
}

func TestDivisiableSumPairs(t *testing.T) {
	tt := []testdata{
		{n: 6, k: 3, ar: []int32{1, 3, 2, 6, 1, 2}, expect: 5},
		{n: 6, k: 5, ar: []int32{1, 2, 3, 4, 5, 6}, expect: 3},
	}

	for _, tc := range tt {
		if divisibleSumPairs(tc.n, tc.k, tc.ar) != tc.expect {
			t.Fail()
		}
	}
}

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	return 0
}
