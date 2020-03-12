package hrk

import (
	"sort"
	"testing"
)

func TestMiniMaxSum(t *testing.T) {
	input := [][]int64{
		[]int64{1, 2, 3, 4, 5},
		[]int64{256741038, 623958417, 467905213, 714532089, 938071625},
	}
	const (
		min int64 = 1 + 2 + 3 + 4
		max int64 = 5 + 2 + 3 + 4
	)
	expect := [][2]int64{
		[2]int64{min, max},
		[2]int64{2063136757, 2744467344},
	}

	for k, v := range input {
		if get := miniMaxSum(v); get != expect[k] {
			t.Errorf("Expect: %v, but: %v", expect, get)
		}
	}
}

func miniMaxSum(arr []int64) [2]int64 {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	return [2]int64{
		sliceSum(arr[:len(arr)-1]),
		sliceSum(arr[1:]),
	}
}

func sliceSum(arr []int64) int64 {
	sum := int64(0)
	for _, v := range arr {
		sum += v
	}
	return sum
}
