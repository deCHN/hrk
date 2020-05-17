package hrk_test

import (
	"strings"
	"testing"
)

func TestAppendAndDelete(t *testing.T) {
	tests := []struct {
		s, t, want string
		k          int32
	}{
		{s: "aba", t: "aba", k: 7, want: "Yes"},
		{s: "hackerhappy", t: "hackerrank", k: 9, want: "Yes"},
		{s: "ashley", t: "ash", k: 2, want: "No"},
		{s: "aaaaaaaaaa", t: "aaaaa", k: 7, want: "Yes"}, //case 2
		{s: "zzzzz", t: "zzzzzzz", k: 4, want: "Yes"},    // case 3
		{s: "y", t: "yu", k: 2, want: "No"},              // case 5
		{s: "asdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcv", t: "asdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcv", k: 20, want: "Yes"}, // case 7
		{s: "abcd", t: "abcdert", k: 10, want: "No"}, // case 10
	}

	for _, v := range tests {
		if get := appendAndDelete(v.s, v.t, v.k); get != v.want {
			t.Errorf("Given %v, want %v, but get %v.\n", v, v.want, get)
		}
	}
}

// https://www.hackerrank.com/challenges/append-and-delete/problem
func appendAndDelete(s string, t string, k int32) string {
	if len(s) >= len(t) {
		if i := strings.Index(s, t); i != -1 {
			if i+len(t) <= int(k) {
				return "Yes"
			}
		}
	}

	diff := len(t)

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

	if (del + app) == int(k) {
		return "Yes"
	} else if (del + app) < int(k) { // 1 < 2
		if len(s)+len(t) > int(k) {
			return "No"
		}
		return "Yes"
	}

	return "No"
}
