package hrk_test

import (
	"sort"
	"testing"
)

func TestHurdleRace(t *testing.T) {
	tests := []struct {
		k      int32
		height []int32
		want   int32
	}{
		{1, []int32{1, 2, 3, 3, 2}, 2},
		{4, []int32{1, 6, 3, 5, 2}, 2},
		{7, []int32{2, 5, 4, 5, 2}, 0},
	}

	for _, v := range tests {
		if get := hurdleRace(v.k, v.height); get != v.want {
			t.Errorf("Give %v, want %v, but %v.\n", v.k, v.want, get)
		}
	}
}

func hurdleRace(k int32, height []int32) int32 {
	sort.Slice(height, func(i int, j int) bool { return height[i] > height[j] })

	var dt int32
	if len(height) > 0 {
		dt = height[0] - k
		if dt <= 0 {
			return 0
		}
		return dt
	}

	return 0
}
