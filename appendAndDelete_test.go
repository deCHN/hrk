package hrk_test

import "testing"

func TestAppendAndDelete(t *testing.T) {
	tests := []struct {
		s, t, want string
		k          int32
	}{
		{s: "aba", t: "aba", k: 7, want: "yes"},
		{s: "hackerhappy", t: "hackerrank", k: 9, want: "yes"},
	}

	for _, v := range tests {
		if get := appendAndDelete(v.s, v.t, v.k); get != v.want {
			t.Errorf("Given %v, want %v, but get %v.\n", v, v.want, get)
		}
	}
}

// https://www.hackerrank.com/challenges/append-and-delete/problem
func appendAndDelete(s string, t string, k int32) string {

	return ""
}
