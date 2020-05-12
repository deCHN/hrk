package hrk_test

import (
	"math/big"
	"testing"
)

func TestExtraLongFactorials(t *testing.T) {
	tests := []struct {
		n    int32
		want string
	}{
		{25, "15511210043330985984000000"},
	}

	for _, v := range tests {
		if get := extraLongFactorials(v.n); get != v.want {
			t.Errorf("Given %v, want %v, but get %v.\n", v.n, v.want, get)
		}
	}
}

// https://www.hackerrank.com/challenges/extra-long-factorials/problem
func extraLongFactorials(n int32) string {
	return big.NewInt(0).MulRange(1, int64(n)).String()
}
