package hrk_test

import "testing"

func TestCatAndMouse(t *testing.T) {
	tests := []struct {
		x, y, z int32
		expect  string
	}{
		{1, 2, 3, "Cat B"},
		{1, 3, 2, "Mouse C"},
	}

	for _, v := range tests {
		if get := catAndMouse(v.x, v.y, v.z); get != v.expect {
			t.Errorf("Given %v, expect %v, but %v.", v, v.expect, get)
		}
	}
}

func catAndMouse(x int32, y int32, z int32) string {
	return "a"
}
