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
		//{[]int32{1, 2, 3, 4, 3, 3, 2, 1}, []int32{8, 6, 4, 1}},
		{[]int32{5, 4, 4, 2, 2, 8}, []int32{6, 4, 2, 1}},
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

	sort.SliceStable(arr, func(i int, j int) bool { return arr[i] < arr[j] })

	l := len(arr)
	s := int32(0)
	r := []int32{}

	/*
		{	2,2,4,4,5,8}
						 1 -> 1
						 1
						 1
					  1 1 -> 2
				 1 1 1 1 -> 4
				 1 1 1 1
			1 1 1 1 1 1 -> 6
			1 1 1 1 1 1
	*/
	for k, v := range arr {
		if s == v {
			continue
		}
		s = v
		r = append(r, int32(l-k))
	}

	return r
}

func CompareSlice(a, b []int32) (bool, error) {
	if len(a) != len(b) {
		return false, errors.New("Different size.")
	}

	for k, v := range a {
		if v != b[k] {
			return false, errors.New(fmt.Sprintf("a[%d] = %v != b[%d] = %v.\n", k, v, k, b[k]))
		}
	}

	return true, nil
}
