package hrk_test

import "testing"

func TestJumpingOnCloudsRevisited(t *testing.T) {
	tests := []struct {
		c       []int32
		k, want int32
	}{
		{[]int32{0, 0, 1, 0, 0, 1, 1, 0}, 2, 92},
	}

	for _, v := range tests {
		if get := jumpingOnCloudsRevisited(v.c, v.k); get != v.want {
			t.Errorf("Given %v, want %v, but get %v.\n", v.c, v.want, get)
		}
	}

}

//https://www.hackerrank.com/challenges/jumping-on-the-clouds-revisited/problem
func jumpingOnCloudsRevisited(c []int32, k int32) int32 {
	n := int32(len(c))
	e := 100

	for i := int32(k); i%n != 0; i = (i + k) % n {
		if c[i] == 1 {
			e = e - 3
		} else {
			e = e - 1
		}
	}

	if c[0] == 0 {
		return int32(e - 1)
	}
	return int32(e - 3)
}
