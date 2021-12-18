package hrk

import "testing"

func TestReverseArray(t *testing.T) {
	td := []struct {
		in, out []int32
	}{
		/* Case 1: */ {[]int32{1, 2, 3}, []int32{3, 2, 1}},
		/* Case 2: */ {[]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 5}, []int32{5, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
	}

	for i, tc := range td {
		got := reverseArray(tc.in)

		// Compare two array
		for k, v := range got {
			if v != tc.out[k] {
				t.Errorf("Testcase %d failed. Expect: %v, but: %v.\n", i, tc.out, got)
			}
		}
	}
}

/*
 * Complete the 'reverseArray' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func reverseArray(a []int32) []int32 {
	l := len(a)
	re := make([]int32, l)

	for k, v := range a {
		re[l-k-1] = v
	}

	return re
}
