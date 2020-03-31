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
		{2100, "13.09.2100"}, //smaple 54
		{2000, "12.09.2000"}, //smaple 58
		{1918, "26.09.1918"}, //smaple 59
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

	//The transition from the Julian to Gregorian calendar system occurred in 1918,
	//when the next day after January 31th was February 14th.
	if year == 1918 {
		monthday[1] = 15
	}

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

// From 1700 to 1917, Russia's official calendar was the Julian calendar;
// since 1919 they used the Gregorian calendar system.
func isLeapYear(year int32) bool {
	if year%400 == 0 {
		return true
	}
	if year < 1919 {
		return year%4 == 0
	} else {
		return year%4 == 0 && year%100 != 0
	}
}
