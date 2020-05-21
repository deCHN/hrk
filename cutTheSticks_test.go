package hrk_test

import (
	"errors"
	"fmt"
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
	return []int32{}
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
