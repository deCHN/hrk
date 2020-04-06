package hrk_test

import "testing"

func TestMoneySpent(t *testing.T) {
	td := []struct {
		keyboards []int32
		drivers   []int32
		b         int32
		expect    int32
	}{
		{[]int32{3, 1}, []int32{5, 2, 8}, 10, 9},
		{[]int32{4}, []int32{5}, 5, -1},
	}

	for _, v := range td {
		if get := getMoneySpent(v.keyboards, v.drivers, v.b); get != v.expect {
			t.Errorf("Given %v, expect %v, but %v.", v, v.expect, get)
		}
	}
}

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	max := int32(-1)

	for _, keyboard := range keyboards {
		for _, drive := range drives {
			if sum := keyboard + drive; sum <= b && sum > max {
				max = sum
			}
		}
	}

	return max
}
