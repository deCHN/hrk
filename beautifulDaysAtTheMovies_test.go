package hrk_test

import "testing"

func TestBeautifulDays(t *testing.T) {
	tests := []struct {
		i, j, k int32
		want    int32
	}{
		{20, 23, 6, 2},
	}

	for _, v := range tests {
		if get := beautifulDays(v.i, v.j, v.k); get != v.want {
			t.Errorf("Give %v, want %v, but get %v.\n", v, v.want, get)
		}
	}
}

func beautifulDays(i int32, j int32, k int32) int32 {
	return 0
}
