package hrk_test

import "testing"

func TestAcmTeam(t *testing.T) {
	tests := []struct {
		topic []string
		want  []int32
	}{
		{[]string{"10101", "11110", "00010"}, []int32{5, 1}},
		{[]string{"10101", "11100", "11010", "00101"}, []int32{5, 2}},
		{[]string{"11101", "10101", "11001", "10111", "10000", "01110"}, []int32{5, 6}}, //Test case 1
	}

	for _, v := range tests {
		if get := acmTeam(v.topic); get != nil {
			if len(get) != len(v.want) {
				t.Errorf("Size not the same. Want %v, but get %v.\n", len(v.want), len(get))
			}
			for k, w := range get {
				if w != v.want[k] {
					t.Errorf("Given %v, want %v, but get %v.\n", v, v.want[k], w)
				}
			}
		} else {
			t.Errorf("Failed. nil returned.\n")
		}
	}
}

// https://www.hackerrank.com/challenges/acm-icpc-team/problem
func acmTeam(topic []string) []int32 {
	max, n := int32(0), int32(0)
	for k, v := range topic {
		for i := k + 1; i < len(topic); i++ {
			set := int32(0)
			for skill := range v {
				if v[skill] == 49 || topic[i][skill] == 49 { // ASCII(49) -> 1
					set++
				}
			}
			if set > max {
				max = set
				n = 0
			}
			if set == max {
				n++
			}

		}
	}

	return []int32{max, n}
}
