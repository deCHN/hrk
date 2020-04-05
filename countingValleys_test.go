package hrk_test

import "testing"

func TestCountingValleys(t *testing.T) {
	td := []struct {
		n      int32
		s      string
		expect int32
	}{
		{8, "UDDDUDUU", 1},
	}

	for _, v := range td {
		if get := countingValleys(v.n, v.s); get != v.expect {
			t.Errorf("Give %v, expect %v, but %v.", v, v.expect, get)
		}
	}
}

func countingValleys(n int32, s string) int32 {
	return 0
}
