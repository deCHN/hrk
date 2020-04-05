package hrk_test

import "testing"

func TestCountingValleys(t *testing.T) {
	td := []struct {
		n      int32
		s      string
		expect int32
	}{
		{8, "UDDDUDUU", 1},
		{10, "UUUDDDUDUU", 0},
		{2, "DU", 1},
		{3, "DUD", 1},
		{4, "DUDU", 2},
		{2, "UD", 0},
		{3, "UDD", 0},
		{4, "UDUD", 0},
	}

	for _, v := range td {
		if get := countingValleys(v.n, v.s); get != v.expect {
			t.Errorf("Give %v, expect %v, but %v.", v, v.expect, get)
		}
	}
}

func countingValleys(n int32, s string) int32 {
	var count, at int32
	inValley := false

	for _, v := range s {
		if v == 'U' {
			at++
		} else {
			at--
		}
		if at == -1 {
			inValley = true
		}
		if at == 0 && inValley {
			count++
			inValley = false
		}
	}

	return count
}
