package hrk_test

import (
	"testing"
)

func TestFindPoint(t *testing.T) {
	tests := []struct {
		px, py, qx, qy int32
		want           []int32
	}{
		{0, 0, 1, 1, []int32{2, 2}},
		{1, 1, 2, 2, []int32{3, 3}},
		{2, 2, 1, 1, []int32{0, 0}},
	}

	for _, v := range tests {
		get := findPoint(v.px, v.py, v.qx, v.qy)

		if len(get) != len(v.want) {
			t.Fatal("Fail. Length doesn't match.")
		}

		for k, item := range v.want {
			if item != get[k] {
				t.Errorf("Give %v, want %v, but get %v.\n", v, item, get[k])
			}
		}
	}
}

func findPoint(px int32, py int32, qx int32, qy int32) []int32 {
	dx := qx - px + qx
	dy := qy - py + qy

	return []int32{dx, dy}
}
