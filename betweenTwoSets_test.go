package hrk

import (
	"sort"
	"testing"
)

func TestGetTotalX(t *testing.T) {

	a := [][]int32{
		[]int32{2, 4},
		[]int32{3, 4},
		[]int32{1},
	}
	b := [][]int32{
		[]int32{16, 32, 96},
		[]int32{24, 48},
		[]int32{100},
	}

	expect := []int32{3, 2, 9}

	for k, exp := range expect {
		if get := getTotalX(a[k], b[k]); get != exp {
			t.Fail()
		}
	}
}

func getTotalX(a []int32, b []int32) int32 {
	//x % a[i] -> 0
	//b[i] % x -> 0
	sort.SliceStable(b, func(i int, j int) bool {
		return b[i] > b[j]
	})

	var n int32
	for x := int32(1); x <= b[0]; x++ {
		if isMultipleOf(x, a) && isFactorOf(x, b) {
			n++
		}
	}

	return n
}

func isMultipleOf(x int32, fc []int32) bool {
	for _, v := range fc {
		if x%v != 0 {
			return false
		}
	}

	return true
}

func isFactorOf(x int32, fc []int32) bool {
	for _, v := range fc {
		if v%x != 0 {
			return false
		}
	}

	return true
}
