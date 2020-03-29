package hrk

import (
	"testing"
)

type testdata struct {
	n, k   int32
	ar     []int32
	expect int32
}

func TestDivisiableSumPairs(t *testing.T) {
	tt := []testdata{
		{n: 5, k: 2, ar: []int32{5, 9, 10, 7, 4}, expect: 4}, //Testcase 2
		{n: 6, k: 3, ar: []int32{1, 3, 2, 6, 1, 2}, expect: 5},
		{n: 6, k: 5, ar: []int32{1, 2, 3, 4, 5, 6}, expect: 3},
		{n: 1, k: 2, ar: []int32{1}, expect: 0},
		{n: 2, k: 2, ar: []int32{1, 1}, expect: 1},
		{n: 3, k: 1, ar: []int32{1, 2, 3}, expect: 3},
		{n: 3, k: 100, ar: []int32{1, 2, 3}, expect: 0},
		{n: 3, k: 100, ar: []int32{100, 100, 100}, expect: 3},
		{n: 100, k: 67, ar: []int32{
			57, 46, 3, 24, 53, 30, 53, 90, 50, 44,
			80, 33, 55, 37, 97, 37, 82, 33, 80, 97,
			84, 18, 85, 28, 99, 77, 93, 90, 88, 46,
			48, 27, 3, 37, 46, 71, 98, 11, 4, 75,
			90, 48, 10, 7, 46, 61, 90, 100, 59, 16,
			27, 84, 18, 59, 51, 74, 52, 40, 12, 25,
			48, 41, 5, 99, 80, 84, 23, 65, 96, 65,
			97, 8, 87, 76, 33, 72, 76, 89, 19, 71,
			39, 6, 33, 74, 55, 26, 6, 98, 80, 56,
			58, 91, 47, 69, 29, 47, 88, 15, 11, 12,
		}, expect: 59}, //Testcase 8
		{n: 100, k: 22, ar: []int32{
			43, 95, 51, 55, 40, 86, 65, 81, 51, 20,
			47, 50, 65, 53, 23, 78, 75, 75, 47, 73,
			25, 27, 14, 8, 26, 58, 95, 28, 3, 23,
			48, 69, 26, 3, 73, 52, 34, 7, 40, 33,
			56, 98, 71, 29, 70, 71, 28, 12, 18, 49,
			19, 25, 2, 18, 15, 41, 51, 42, 46, 19,
			98, 56, 54, 98, 72, 25, 16, 49, 34, 99,
			48, 93, 64, 44, 50, 91, 44, 17, 63, 27,
			3, 65, 75, 19, 68, 30, 43, 37, 72, 54,
			82, 92, 37, 52, 72, 62, 3, 88, 82, 71,
		}, expect: 216}, //Testcase 6
	}

	for _, tc := range tt {
		if get := divisibleSumPairs(tc.n, tc.k, tc.ar); get != tc.expect {
			t.Errorf("Expect %v, but %v.\n", tc.expect, get)
		}
	}
}

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	pairs := 0

	for index, v := range ar {
		for scan := int32(index + 1); scan < n; scan++ {
			if (v+ar[scan])%k == 0 {
				//fmt.Println(v, "+", ar[scan], "%", k)
				pairs++
			}
		}
	}

	return int32(pairs)
}
