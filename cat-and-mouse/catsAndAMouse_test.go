package hrk_test

import (
	"math"
	"sort"
	"testing"
)

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
	if x == z {
		return "Cat A"
	}
	if y == z {
		return "Cat B"
	}
	if x == z && y == z {
		return "Mouse C"
	}

	p := [3]int32{x, y, z}
	sort.Slice(p[:], func(i int, j int) bool { return p[i] < p[j] })

	if p[1] == z {
		a := math.Abs(float64(x - z))
		b := math.Abs(float64(y - z))
		if a == b {
			return "Mouse C"
		}
		if a < b {
			return "Cat A"
		}
		return "Cat B"
	}

	if p[1] == x {
		return "Cat A"
	}

	return "Cat B"
}
