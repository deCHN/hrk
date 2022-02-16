package hrk

import "testing"

/*
 *Given an array  of  distinct elements. Let  and  be the smallest and the next smallest element in the interval  where .
 *
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
			t.Errorf("Case %d -\t want %v, but get %v.", i+1, tc.want, get)
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
	return 0
}
