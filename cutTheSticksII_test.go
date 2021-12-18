package hrk

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestCTS(t *testing.T) {
	td := []struct {
		in, out []int
	}{
		/* case 1 */ {[]int{1}, []int{1}},
		/* case 2 */ {[]int{1, 1}, []int{2}},
		/* case 3 */ {[]int{1, 2, 3, 4, 3, 3, 2, 1}, []int{8, 6, 4, 1}},
		/* case 4 */ {[]int{2, 2, 3, 4}, []int{4, 2, 1}},
		/* case 5 */ {[]int{5, 4, 4, 2, 2, 8}, []int{6, 4, 2, 1}},
	}

	for i, d := range td {
		if get := cuts(d.in); !reflect.DeepEqual(get, d.out) { // check standard lib for slice comparation
			fmt.Printf("Testcase %d failed. Expect: %v, but: %v.\n", i+1, d.out, get)
		}
	}
}

// https://www.hackerrank.com/challenges/cut-the-sticks/problem
func cuts(arr []int) []int {
	cuts := make([]int, 0)
	sort.SliceStable(arr, func(i, j int) bool { return arr[i] < arr[j] })

	for len(arr) > 0 {
		cuts = append(cuts, len(arr))

		after := make([]int, 0)

		for _, v := range arr {
			if v != arr[0] {
				after = append(after, v-arr[0])
			}
		}
		arr = after
	}
	return cuts
}
