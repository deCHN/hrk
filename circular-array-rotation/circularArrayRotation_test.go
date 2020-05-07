package hrk_test

import "testing"

func TestCircularArrayRotation(t *testing.T) {
	tests := []struct {
		a       []int32
		k       int32
		queries []int32
		want    []int32
	}{
		{[]int32{1, 2, 3}, 2, []int32{0, 1, 2}, []int32{2, 3, 1}},
	}

	for _, v := range tests {
		get := circularArrayRotation(v.a, v.k, v.queries)

		if len(get) != len(v.want) {
			t.Fatal("Test fail. Length doesn't match.")
		}

		for k, i := range v.want {
			if i != get[k] {
				t.Errorf("Given %v, want %v, but get %v.\n", v.a, i, get[k])
			}
		}
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
	var ca circularArray = a
	rotate := ca.rotate(k)

	var out []int32
	for _, v := range queries {
		out = append(out, rotate[v])
	}

	return out
}

type circularArray []int32

func (a circularArray) rotate(k int32) []int32 {
	size := len(a)
	for i := 0; i < int(k); i++ {
		var t []int32
		last := a[size-1]
		t = append(t, last)
		a = a[:size-1]
		t = append(t, a...)
		a = t
	}

	return a
}
