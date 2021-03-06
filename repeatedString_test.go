package hrk_test

import (
	"math"
	"strings"
	"testing"
)

func TestRepeatedString(t *testing.T) {
	tests := []struct {
		s       string
		n, want int64
	}{
		{"aba", 10, 7},
		{"a", 100000, 100000},
	}

	for _, v := range tests {
		if get := repeatedString(v.s, v.n); get != v.want {
			t.Errorf("Given %v, want %v, but get %v.\n", v, v.want, get)
		}
	}
}

// https://www.hackerrank.com/challenges/repeated-string/problem
// 1 <= |s| <= 100
// 1 <= n <= 10^12
func repeatedString(s string, n int64) int64 {
	sz := int64(len(s))

	ns := int64(math.Floor(float64(n / sz)))

	a := int64(strings.Count(s, "a"))

	na := a * ns

	as := n % sz

	for i := int64(0); i < as; i++ {
		if s[i] == 'a' {
			na++
		}
	}

	return na
}
