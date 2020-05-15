package hrk_test

import (
	"testing"
)

func TestAppendAndDelete(t *testing.T) {
	tests := []struct {
		s, t, want string
		k          int32
	}{
		//{s: "aba", t: "aba", k: 7, want: "Yes"},
		//{s: "hackerhappy", t: "hackerrank", k: 9, want: "Yes"},
		//{s: "ashley", t: "ash", k: 2, want: "No"},
		{s: "aaaaaaaaaa", t: "aaaaa", k: 7, want: "Yes"},
	}

	for _, v := range tests {
		if get := appendAndDelete(v.s, v.t, v.k); get != v.want {
			t.Errorf("Given %v, want %v, but get %v.\n", v, v.want, get)
		}
	}
}

// https://www.hackerrank.com/challenges/append-and-delete/problem
func appendAndDelete(s string, t string, k int32) string {
	var diff int

	// s: abci -> len(s) - k
	// t: abcde  -> k = 3 -> len(t) - k
	for k, v := range t {
		if k == len(s) {
			diff = k
			break
		}
		if byte(v) != s[k] {
			diff = k
			break
		}
	}

	del := len(s) - diff
	app := len(t) - diff

	//fmt.Printf("Del: %v, Append: %v.\n", del, app)

	if (del + app) <= int(k) {
		return "Yes"
	}

	return "No"
}
