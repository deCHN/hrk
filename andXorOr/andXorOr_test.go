package hrk

import (
	"fmt"
	"testing"
)

/*
 *Given an array A[] of N distinct elements. Let M1 and M2 be the smallest and the next smallest element in the interval [L, R] where 1 <= L < R <= N.
 *.
 *
 *where , are the bitwise operators ,  and  respectively.
 *Your task is to find the maximum possible value of .
 *
 *Input Format
 *
 *First line contains integer .
 *Second line contains  integers, representing elements of the array .
 */
func TestAndXorOr(t *testing.T) {
	td := []struct {
		a    []int32
		want int32
	}{
		{[]int32{9, 6, 4, 5, 2}, 15},
	}

	for i, tc := range td {
		get := andXorOr(tc.a)

		if get != tc.want {
			t.Errorf("Case %d failed. Want %v, but get %v.", i+1, tc.want, get)
		}
	}
}

/*
 * Complete the 'andXorOr' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */
func andXorOr(a []int32) int32 {
	o := 9 & 6
	p := 9 | 6
	q := 9 ^ 6
	fmt.Println("9 and 6", o)
	fmt.Println("9 or 6", p)
	fmt.Println("9 xor 6", q)
	fmt.Println("o xor p and q", o^p&q)

	return 0
}
