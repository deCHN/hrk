package hrk_test

import "testing"

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

func repeatedString(s string, n int64) int64 {
	return 0
}
