package hrk_test

import "testing"

func TestAcmTeam(t *testing.T) {
	tests := []struct {
		topic []string
		want  []int32
	}{
		{[]string{"10101", "11110", "00010"}, []int32{5, 1}},
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
	return nil
}
