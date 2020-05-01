package hrk_test

import "testing"

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
			t.Errorf("Given %v, want %v, but %v./n", v, v.want, get)
		}
	}
}

func saveThePrisoner(n int32, m int32, s int32) int32 {
	return 0
}
