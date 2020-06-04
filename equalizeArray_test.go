package hrk_test

import "testing"

func TestEqualizeArray(t *testing.T) {
	tests := []struct {
		arr  []int32
		want int32
	}{
		{[]int32{1, 2, 2, 3}, 2},
		{[]int32{3, 3, 2, 1, 3}, 2},
	}

	for _, v := range tests {
		if get := equalizeArray(v.arr); get != v.want {
			t.Errorf("Given %v, want %v, but get %v.\n", v, v.want, get)
		}
	}
}

func equalizeArray(arr []int32) int32 {
	times := make(map[int32]int)

	for _, v := range arr {
		times[v]++
	}

	max := 0
	for _, v := range times {
		if v > max {
			max = v
		}
	}

	return int32(len(arr) - max)
}
