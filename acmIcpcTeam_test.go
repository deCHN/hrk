package hrk_test

import "testing"

func TestAcmTeam(t *testing.T) {
	tests := []struct {
		topic []string
		want  []int32
	}{
		{[]string{"10101", "11110", "00010"}, []int32{5, 1}},
	}
}

// https://www.hackerrank.com/challenges/acm-icpc-team/problem
func acmTeam(topic []string) []int32 {

}
