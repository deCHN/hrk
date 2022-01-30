package hrk

import "testing"

/*
 *you have three stacks of cylinders where each cylinder has the same diameter, but they may vary in height. You can change the height of a stack by removing and discarding its topmost cylinder any number of times.
 *
 *Find the maximum possible height of the stacks such that all of the stacks are exactly the same height. This means you must remove zero or more cylinders from the top of zero or more of the three stacks until they are all the same height, then return the height.
 *
 *Example
 *
 *
 *
 *
 *There are  and  cylinders in the three stacks, with their heights in the three arrays. Remove the top 2 cylinders from  (heights = [1, 2]) and from  (heights = [1, 1]) so that the three stacks all are 2 units tall. Return  as the answer.
 *
 *Note: An empty stack is still a stack.
 *
 *Function Description
 *
 *Complete the equalStacks function in the editor below.
 *
 *equalStacks has the following parameters:
 *
 *int h1[n1]: the first array of heights
 *int h2[n2]: the second array of heights
 *int h3[n3]: the third array of heights
 *Returns
 *
 *int: the height of the stacks when they are equalized
 *Input Format
 *
 *The first line contains three space-separated integers, , , and , the numbers of cylinders in stacks , , and . The subsequent lines describe the respective heights of each cylinder in a stack from top to bottom:
 *
 *The second line contains  space-separated integers, the cylinder heights in stack . The first element is the top cylinder of the stack.
 *The third line contains  space-separated integers, the cylinder heights in stack . The first element is the top cylinder of the stack.
 *The fourth line contains  space-separated integers, the cylinder heights in stack . The first element is the top cylinder of the stack.
 *Constraints
 */
func TestEqualStacks(t *testing.T) {
	td := []struct {
		h1   []int32
		h2   []int32
		h3   []int32
		want int32
	}{
		{[]int32{3, 2, 1, 1, 1}, []int32{4, 3, 2}, []int32{1, 1, 4, 1}, 5},
	}

	for k, tc := range td {
		get := equalStacks(tc.h1, tc.h2, tc.h3)

		if get != tc.want {
			t.Errorf("Case %d failed. Want %v, but get %v.\n", k, tc.want, get)
		}
	}
}

/*
 * Complete the 'equalStacks' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY h1
 *  2. INTEGER_ARRAY h2
 *  3. INTEGER_ARRAY h3
 */
func equalStacks(h1 []int32, h2 []int32, h3 []int32) int32 {
	stc1, t1 := hstack(h1)
	stc2, t2 := hstack(h2)
	stc3, t3 := hstack(h3)

	for k, v := range stc1 {

	}

	return 0
}

func hstack(h []int32) (map[int32]bool, int32) {
	var height int32

	for _, v := range h {
		height += v
	}

	r := height

	hs := make(map[int32]bool)
	hs[height] = true

	for _, v := range h {
		hs[height-v] = true
		height = height - v
	}

	return hs, r
}
