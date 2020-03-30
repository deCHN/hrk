package hrk

import (
	"sort"
	"testing"
)

type testData struct {
	arr    []int32
	expect int32
}

func TestMigratoryBirds(t *testing.T) {
	input := []testData{
		testData{arr: []int32{1, 4, 4, 4, 5, 3}, expect: 4},
		testData{arr: []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}, expect: 3},
	}

	for _, td := range input {
		if get := migratoryBirds(td.arr); get != td.expect {
			t.Errorf("Expect %v, but %v.\n", td.expect, get)
		}
	}
}

func migratoryBirds(arr []int32) int32 {
	signing := make(map[int32]int)

	for _, v := range arr {
		signing[v]++
	}

	//Constraints type is max 5
	values := make([]int, 5)

	for _, v := range signing {
		values = append(values, v)
	}

	sort.SliceStable(values, func(i int, j int) bool {
		return values[i] > values[j]
	})

	var most int
	if len(values) > 0 {
		most = values[0]
	}

	bird := int32(100)
	for k, v := range signing {
		if v == most && k < bird {
			bird = k
		}
	}

	return bird
}
