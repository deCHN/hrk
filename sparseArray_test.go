package hrk

import "testing"

/*
 * There is a collection of input strings and a collection of query strings. For each query string, determine how many times it occurs in the list of input strings. Return an array of the results.
 */

func TestMatchingStrings(t *testing.T) {
	tcs := []struct {
		in   [2][]string
		want []int32
	}{
		{[2][]string{[]string{"aba", "baba", "aba", "xzxb"}, []string{"aba", "xzxb", "ab"}}, []int32{2, 1, 0}},
		{[2][]string{[]string{"def", "de", "fgh"}, []string{"de", "lmn", "fgh"}}, []int32{1, 0, 1}},
	}

	for _, tc := range tcs {
		get := matchingStrings(tc.in[0], tc.in[1])

		for k, v := range tc.want {
			if get[k] != v {
				t.Errorf("Want: %v, but get: %v.", v, get[k])
			}
		}
	}

}

func matchingStrings(strings []string, queries []string) []int32 {
	r := make([]int32, len(queries))

	for k, q := range queries {
		for _, v := range strings {
			if v == q {
				r[k]++
			}
		}
	}

	return r
}
