package hrk_test

import "testing"

func TestLibraryFine(t *testing.T) {
	tests := []struct {
		d1, m1, y1, d2, m2, y2, want int32
	}{
		{9, 6, 2015, 6, 6, 2015, 45},
	}

	for _, v := range tests {
		if get := libraryFine(v.d1, v.m1, v.y1, v.d2, v.m2, v.y2); get != v.want {
			t.Errorf("Given %v, want %v, but get %v.\n", v, v.want, get)
		}
	}

}

// https://www.hackerrank.com/challenges/library-fine/problem
// d1, m1, y1: returned date day, month and year
// d2, m2, y2: due date day, month and year
func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
	if y1-y2 > 0 {
		return 10000
	}

	if m := m1 - m2; m > 0 {
		return m * 500
	}

	return (d1 - d2) * 15
}
