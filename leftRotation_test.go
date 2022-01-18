package hrk

import "testing"

func TestLeftRotation(t *testing.T) {
	td := []struct {
		d      int32
		input  []int32
		expect []int32
	}{
		{4, []int32{1, 2, 3, 4, 5}, []int32{5, 1, 2, 3, 4}},
		{2, []int32{4, 5}, []int32{4, 5}},
		{1, []int32{10, 11}, []int32{11, 10}},
		{0, []int32{10, 11}, []int32{10, 11}},
	}

	for _, v := range td {
		get := rotateLeft(v.d, v.input)
		if len(get) != len(v.expect) {
			t.Errorf("Array size doesn't match. Expect %v, but %v.", len(v.expect), len(get))
		}
		for i, e := range get {
			if e != v.expect[i] {
				t.Errorf("Array element doesn't match. Expect: %v, but: %v", v.expect[i], e)
			}
		}
	}
}

func rotateLeft(d int32, in []int32) []int32 {
	return append(in[d:], in[:d]...)
}
