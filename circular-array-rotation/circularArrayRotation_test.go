package hrk_test

import "testing"

func TestCircularArrayRotation(t *testing.T) {
	tests := []struct {
		a, k, queries []int32
		want          []int32
	}{
		{[]int32{3, 2, 3}, []int32{1, 2, 3}, []int32{0, 1, 2}, []int32{2, 3, 1}},
	}
}

/*
 * https://www.hackerrank.com/challenges/circular-array-rotation/problem
 * a: 包含了3个空格分隔的整数N, K 和 Q。
 * k: 包含了N个空格分隔的整数，表示最初数组A的元素。
 * queries: 接下来的Q行，每行包含一个整数X。
 *
 * Constraints:
 * - 1 <= N <= 10**5
 * - 1 <= A[i] <= 10**5
 * - 1 <= K <= 10**5
 * - 1 <= Q <= 500
 * - 0 <= X <= N-1
 */
func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	return nil
}
