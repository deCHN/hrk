package hrk_test

import "testing"

func TestCutTheSticks(t *testing.T) {
	tests := []struct {
		arr, want []int32
	}{
		{[]int32{1, 2, 3, 4, 3, 3, 2, 1}, []int32{8, 6, 4, 1}},
	}

	for _, v := range tests {
		get := cutTheSticks(v.arr)
		if ok, err := compareSlice(v.want, get); !ok || err != nil {
			t.Errorf("Given %v, want %v, but get %v.\n", v.arr, v.want, get)
		}
	}
}

// https://www.hackerrank.com/challenges/cut-the-sticks/problem
func cutTheSticks(arr []int32) []int32 {
	return []int32{}
}

func compareSlice(a, b []int32) (bool, error) {

}
