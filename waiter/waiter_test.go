package main

import (
	"fmt"
	"testing"
)

/*
 * https://www.hackerrank.com/challenges/waiter/problem?utm_campaign=challenge-recommendation&utm_medium=email&utm_source=24-hour-campaign
 *
 * You are a waiter at a party. There is a pile of numbered plates. Create an empty  array. At each iteration, , remove each plate from the top of the stack in order. Determine if the number on the plate is evenly divisible by the  prime number. If it is, stack it in pile . Otherwise, stack it in stack . Store the values in  from top to bottom in . In the next iteration, do the same with the values in stack . Once the required number of iterations is complete, store the remaining values in  in , again from top to bottom. Return the  array.
 *
 *Example
 *
 *
 *
 *An abbreviated list of primes is . Stack the plates in reverse order.
 *
 *
 *
 *Begin iterations. On the first iteration, check if items are divisible by .
 *
 *
 *Move  elements to .
 *
 *
 *On the second iteration, test if  elements are divisible by .
 *
 *
 *Move  elmements to .
 *
 *
 *And on the third iteration, test if  elements are divisible by .
 *
 *
 *Move  elmements to .
 *
 *
 *All iterations are complete, so move the remaining elements in , from top to bottom, to .
 *
 *. Return this list.
 *
 *Function Description
 *
 *Complete the waiter function in the editor below.
 *
 *waiter has the following parameters:
 *
 *int number[n]: the numbers on the plates
 *int q: the number of iterations
 *Returns
 *
 *int[n]: the numbers on the plates after processing
 *Input Format
 *
 *The first line contains two space separated integers,  and .
 *The next line contains  space separated integers representing the initial pile of plates, i.e., .
 */
func TestWaiter(t *testing.T) {
	td := []struct {
		number []int32
		q      int32
		want   []int32
	}{
		/* 1 */ {[]int32{4, 4, 9, 3, 3}, 2, []int32{4, 4, 3, 3, 9}},
		/* 2 */ {[]int32{3, 4, 7, 6, 5}, 1, []int32{4, 6, 3, 7, 5}},
		/* 3 */ {[]int32{2, 3, 4, 5, 6, 7}, 3, []int32{2, 4, 6, 3, 5, 7}},
		/* 4 */ {[]int32{3, 3, 4, 4, 9}, 2, []int32{4, 4, 9, 3, 3}},
	}

	for k, tc := range td {
		get := waiter(tc.number, tc.q)
		fmt.Println("GET:", get)
		for i, v := range get {
			if v != tc.want[i] {
				t.Errorf("Case %d failed. Want %v, but get %v.", k, tc.want[i], v)
			}
		}
	}
}

/*
 * Complete the 'waiter' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY number
 *  2. INTEGER q
 */
func waiter(number []int32, q int32) []int32 {
	answer := make([]int32, 0)
	prm := []int32{2, 3, 5, 7, 11, 13, 17, 19, 23}
	a := number

	for i := int32(0); i < q; i++ {
		number = a
		a = make([]int32, 0)

		for _, v := range number {
			if v%prm[i] == 0 {
				answer = append(answer, v)
			} else {
				a = append([]int32{v}, a...)
			}
		}
	}

	for i := len(a) - 1; i >= 0; i-- {
		answer = append(answer, a[i])

	}
	return answer
}
