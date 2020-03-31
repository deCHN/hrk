package hrk

import (
	"fmt"
	"testing"
)

func TestDayOfProgrammer(t *testing.T) {
	testdata := []struct {
		year   int32
		expect string
	}{
		{2017, "13.09.2017"}, //sample 0
		{2016, "12.09.2016"}, //sample 1
		{1800, "12.09.1800"}, //smaple 2
		{2100, "13.09.2100"}, //smaple 2
	}

	for _, v := range testdata {
		if get := dayOfProgrammer(v.year); get != v.expect {
			t.Errorf("Expect: %s, but %s.", v.expect, get)
		}
	}
}

// Complete the dayOfProgrammer function below.
func dayOfProgrammer(year int32) string {
	monthday := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if isLeapYear(year) {
		monthday[1] = 29
	}

	DOP := 256

	month := 0
	day := 0
	for k, v := range monthday {
		if DOP-v < 0 {
			month = k + 1
			day = DOP
			break
		}
		DOP = DOP - v
	}

	return fmt.Sprintf("%02d.%02d.%d", day, month, year)
}

func isLeapYear(year int32) bool {
	return year%4 == 0
}
