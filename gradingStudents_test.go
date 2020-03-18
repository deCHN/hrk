package hrk

import "testing"

func TestGradingStudents(t *testing.T) {
	input := []int32{73, 67, 38, 33, 37}
	output := []int32{75, 67, 40, 33, 37}

	get := gradingStudents(input)

	for k, v := range get {
		if v != output[k] {
			t.Errorf("Expect %v, but %v.\n", output[k], v)
		}
	}
}

func gradingStudents(grades []int32) []int32 {
	rt := make([]int32, len(grades))
	for k, v := range grades {
		switch {
		case v < 38:
			rt[k] = v
		case (v % 5) > 2:
			rt[k] = (v + 5) - (v % 5)
		default:
			rt[k] = v
		}
	}

	return rt
}
