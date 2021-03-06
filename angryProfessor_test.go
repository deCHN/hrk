package hrk_test

import (
	"sort"
	"testing"
)

func TestAngryProfessor(t *testing.T) {
	tests := []struct {
		k    int32
		a    []int32
		want string
	}{
		{3, []int32{-1, -3, 4, 2}, "YES"},
		{2, []int32{0, -1, 2, 1}, "NO"},
	}

	for _, v := range tests {
		if get := angryProfessor(v.k, v.a); get != v.want {
			t.Errorf("Given %v, want %v, but %v. \n", v.a, v.want, get)
		}
	}
}

func angryProfessor(k int32, a []int32) string {
	sort.Slice(a, func(i int, j int) bool { return a[i] < a[j] })

	for _, v := range a {
		if k == 0 {
			return "NO"
		}
		if v > 0 {
			return "YES"
		}
		k--
	}
	return "YES"
}
