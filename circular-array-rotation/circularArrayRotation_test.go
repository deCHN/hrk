package hrk_test

import "testing"

func TestCircularArrayRotation(t *testing.T) {
	tests := []struct {
		a, k, queries []int32
		want          []int32
	}{
		{[]int32{3, 2, 3}, []int32{1, 2, 3}, []int32{0, 1, 2}, []int32{2, 3, 1}},
	}
}

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	return nil
}
