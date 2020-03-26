package hrk

import "testing"

type BirthdayTestData struct {
	squares []int32
	d, m    int32
}

func TestBirthday(t *testing.T) {

	input := []BirthdayTestData{
		BirthdayTestData{squares: []int32{1, 2, 1, 3, 2}, d: 3, m: 2},
		BirthdayTestData{squares: []int32{1, 1, 1, 1, 1, 1}, d: 3, m: 2},
		BirthdayTestData{squares: []int32{4}, d: 4, m: 1},
	}

	output := []int32{2, 0, 1}

	for k, v := range input {
		if birthday(v.squares, v.d, v.m) != output[k] {
			t.Fail()
		}
	}
}

func birthday(s []int32, d int32, m int32) int32 {

	for k, v := s {
		
	}
	return 0
}

func cut(s []int32, d int32, m int32, n int32, sum int32) bool{

	if len(s) = 0 {
		return false
	}

	if n > m {
		return  false
	}

	for k, v := range s {
		days := sum + v

		if days > d {
			return false
		}

		if days < d && n < m {
			si := append([]int32(nil), s[:k], s[k+1:])
			cut(si, d, m ,n++, days)
		}
		if days < d && n == m {
			return false
		}
		if days == d && n == m {
			return true
		}
		if n < m {
		}
	}
}
