package hrk

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

/*
 * https://www.hackerrank.com/challenges/maximum-element/problem?utm_campaign=challenge-recommendation&utm_medium=email&utm_source=24-hour-campaign
 * You have an empty sequence, and you will be given  queries. Each query is one of these three types:
 *
 * 1 x  -Push the element x into the stack.
 * 2    -Delete the element present at the top of the stack.
 * 3    -Print the maximum element in the stack.
 * Function Description
 *
 * Complete the getMax function in the editor below.
 *
 * getMax has the following parameters:
 * - string operations[n]: operations as strings
 *
 * Returns
 * - int[]: the answers to each type 3 query
 *
 * Input Format
 *
 * The first line of input contains an integer, . The next  lines each contain an above mentioned query.
 *
 * Constraints
 *
 * Constraints
 *
 *
 *
 * All queries are valid.
 */
func TestGetMax(t *testing.T) {
	td := []struct {
		in   []string
		want []int32
	}{
		{[]string{"1 92", "2", "1 20", "2", "1 26", "1 20", "2", "3", "1 91", "3"}, []int32{26, 91}},
	}

	for _, tc := range td {
		get := getMax(tc.in)

		for i, v := range get {
			if v != tc.want[i] {
				t.Errorf("Element[%d] is %v, but want %v.\n", i, v, tc.want[i])
			}
		}
	}
}

/*
 * Complete the 'getMax' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts STRING_ARRAY operations as parameter.
 */
func getMax(operations []string) []int32 {
	lifo := make([]int32, 0)
	r := make([]int32, 0)

	for _, op := range operations {
		v := strings.Split(op, " ")

		switch v[0] {
		case "1":
			s, err := strconv.Atoi(v[1])
			if err != nil {
				fmt.Errorf("Cannot convert %v to integer.", v[1])
				return nil
			}
			lifo = append(lifo, int32(s))
		case "2":
			lifo = lifo[:len(lifo)-1]
		case "3":
			max := int32(0)
			for _, n := range lifo {
				if n > max {
					max = n
				}
			}
			r = append(r, max)
		}
	}

	return r
}
