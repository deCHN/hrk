package hrk_test

import (
	"container/ring"
	"testing"
)

func TestSaveThePrisoner(t *testing.T) {
	tests := []struct {
		n, m, s int32
		want    int32
	}{
		{7, 19, 2, 6},
		{3, 7, 3, 3},
	}

	for _, v := range tests {
		if get := saveThePrisoner(v.n, v.m, v.s); get != v.want {
			t.Errorf("Given %v, want %v, but %v.\n", v, v.want, get)
		}
	}
}

/*
 * https://www.hackerrank.com/challenges/save-the-prisoner/problem
 *
 * n: the number of prisoners. 1 <=n <= 10^9
 * m: the number of sweets. 1 <=m <= 10^9
 * s: the chair number to start passing out treats at. 1 <= s <= n
 */
func saveThePrisoner(n int32, m int32, s int32) int32 {
	r := ring.New(int(n))

	for i := int32(1); i <= n; i++ {
		r.Value = i
		r = r.Next()
	}

	for i := 0; i < r.Len(); i++ {
		if r.Value == s {
			break
		}
		r = r.Next()
	}

	r = r.Move(int(m - 1))

	return r.Value.(int32)
}
