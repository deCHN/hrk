package hrk_test

import "testing"

func TestJumpingOnClouds(t *testing.T) {
	tests := []struct {
		c    []int32
		want int32
	}{
		{[]int32{0, 0, 1, 0, 0, 1, 0}, 4},
		{[]int32{0, 0, 0, 0, 1, 0}, 3},
	}

	for _, v := range tests {
		if get := jumpingOnClouds(v.c); get != v.want {
			t.Errorf("Given %v, want %v, but get %v.\n", v, v.want, get)
		}
	}

}

func jumpingOnClouds(c []int32) int32 {
	return 0
}
