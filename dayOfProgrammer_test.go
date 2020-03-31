package hrk

import "testing"

func TestDayOfProgrammer(t *testing.T) {
	testdata := []struct {
		year   int32
		expect string
	}{
		{2017, "13.09.2017"}, //sample 0
		{2016, "12.09.2016"}, //sample 1
		{1800, "12.09.1800"}, //smaple 2
	}

	for _, v := range testdata {
		if get := dayOfProgrammer(v.year); get != v.expect {
			t.Errorf("Expect: %s, but %s.", v.expect, get)
		}
	}
}

// Complete the dayOfProgrammer function below.
func dayOfProgrammer(year int32) string {

	return ""
}
