package hrk

import "testing"

func TestKangaroo(t *testing.T) {
	input := [][4]int32{
		[4]int32{0, 3, 4, 2}, //true
		[4]int32{0, 2, 5, 3}, //false
	}

	expect := []bool{true, false}

	for k, v := range input {
		if get := kangaroo(v[0], v[1], v[2], v[3]); get != expect[k] {
			t.Fail()
		}
	}
}

// Original problem is to return "YES" or "NO" string
func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) bool {
	return isPositiveInt(float32(x2-x1) / float32(v1-v2))
}

func isPositiveInt(v float32) bool {
	if v == float32(int(v)) && v > 0 {
		return true
	}
	return false
}
