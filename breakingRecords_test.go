package hrk

import "testing"

func TestBreakingRecords(t *testing.T) {
	input := [][]int32{
		[]int32{10, 5, 20, 20, 4, 5, 2, 25, 1},
		[]int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42},
	}

	expect := [][2]int32{
		[2]int32{2, 4},
		[2]int32{4, 0},
	}

	for k, v := range input {
		if breakingRecords(v) != expect[k] {
			t.Fail()
		}
	}
}

func breakingRecords(scores []int32) [2]int32 {
	min := scores[0]
	max := scores[0]

	var high, low int32

	for _, v := range scores {
		if v < min {
			low++
			min = v
		}
		if v > max {
			high++
			max = v
		}
	}

	return [2]int32{high, low}
}
