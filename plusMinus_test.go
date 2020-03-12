package hrk

import (
	"testing"
)

func TestPlusMinus(t *testing.T) {
	input := []int32{
		-4, 3, -9, 0, 4, 1,
	}

	expect := []float32{
		0.500000,
		0.33333334,
		0.16666667,
	}

	for k, v := range plusMinus(input) {
		if v != expect[k] {
			t.Errorf("Expect %v, but %v.", expect[k], v)
		}
	}
}

// Complete the plusMinus function below.
func plusMinus(arr []int32) []float32 {

	var pos, neg, zero float32

	for _, v := range arr {
		if v == 0 {
			zero++
		} else if v > 0 {
			pos++
		} else {
			neg++
		}
	}
	l := float32(len(arr))

	return []float32{pos / l, neg / l, zero / l}
}
