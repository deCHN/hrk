package hrk_test

import (
	"errors"
	"fmt"
	"sort"
	"testing"
)

func TestCutTheSticks(t *testing.T) {
	tests := []struct {
		arr, want []int32
	}{
		{[]int32{1, 2, 3, 4, 3, 3, 2, 1}, []int32{8, 6, 4, 1}},
	}

	for _, v := range tests {
		get := cutTheSticks(v.arr)
		if ok, err := CompareSlice(v.want, get); !ok || err != nil {
			t.Errorf("Given %v, want %v, but get %v. %v\n", v.arr, v.want, get, err)
		}
	}
}

// https://www.hackerrank.com/challenges/cut-the-sticks/problem
func cutTheSticks(arr []int32) []int32 {

	// a map of stick length to the amount of the sticks with that length
	ln := make(map[int32]int32)

	for _, v := range arr {
		ln[v]++
	}

	sort.SliceStable(arr, func(i int, j int) bool { return arr[i] > arr[j] })
	longest := arr[0]

	r := []int32{}
	l := int32(len(arr))
	for i := int32(1); i <= longest; i++ {
		r = append(r, l-ln[i])
	}

	return r
}

func CompareSlice(a, b []int32) (bool, error) {
	if len(a) != len(b) {
		return false, errors.New("Different size.")
	}

	for k, v := range a {
		if v != b[k] {
			return false, errors.New(fmt.Sprintf("The %d-th element of slice a is %v, slice b is %v.\n", k, v, b[k]))
		}
	}

	return true, nil
}
