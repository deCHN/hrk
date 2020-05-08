package hrk_test

import "testing"

func TestJumpingOnClouds(t *testing.T) {
	tests := []struct {
		c       []int32
		k, want int32
	}{
		{[]int32{0, 0, 1, 0, 0, 1, 1, 0}, 2, 92},
	}

	for _, v := range tests {
		if get := jumpingOnClouds(v.c, v.k); get != v.want {
			t.Errorf("Given %v, want %v, but get %v.\n", v.c, v.want, get)
		}
	}

}

//https://www.hackerrank.com/challenges/jumping-on-the-clouds-revisited/problem
func jumpingOnClouds(c []int32, k int32) int32 {
	return 0
}
