package hrk_test

import (
	"math"
	"testing"
)

func TestBeautifulDays(t *testing.T) {
	tests := []struct {
		i, j, k int32
		want    int32
	}{
		{20, 23, 6, 2}, // (20 - 2) / 6, (22 - 22 ) / 6
	}

	for _, v := range tests {
		if get := beautifulDays(v.i, v.j, v.k); get != v.want {
			t.Errorf("Give %v, want %v, but get %v.\n", v, v.want, get)
		}
	}
}

func beautifulDays(i int32, j int32, k int32) int32 {

	reverse := func(x int32) int32 {

		return 0
	}

	count := 0

	for v := i; v < j; v++ {
		df := v - reverse(v)
		df = int32(math.Abs(float64(df)))

		if math.Mod(float64(df), float64(j)) == 0 {
			count++
		}

	}
	return int32(count)
}
