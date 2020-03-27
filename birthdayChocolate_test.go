package hrk

import (
	"testing"
)

type BirthdayChocolateTestData struct {
	squares []int32
	d, m    int32
}

func TestBirthdayChocolate(t *testing.T) {

	input := []BirthdayChocolateTestData{
		BirthdayChocolateTestData{squares: []int32{2, 2, 1, 3, 2}, d: 4, m: 2},
		BirthdayChocolateTestData{squares: []int32{1, 2, 1, 3, 2}, d: 3, m: 2},
		BirthdayChocolateTestData{squares: []int32{1, 1, 1, 1, 1, 1}, d: 3, m: 2},
		BirthdayChocolateTestData{squares: []int32{4}, d: 4, m: 1},
	}

	output := []int32{2, 2, 0, 1}

	for k, v := range input {
		if birthdayChocolate(v.squares, v.d, v.m) != output[k] {
			t.Fail()
		}
	}
}

func birthdayChocolate(s []int32, d int32, m int32) int32 {
	success := 0

	for k := range s {
		sum := 0
		for i := 0; i < int(m); i++ {
			if k+i >= len(s) {
				break
			}
			sum += int(s[k+i])
		}
		if sum == int(d) {
			success++
		}
	}
	return int32(success)
}
