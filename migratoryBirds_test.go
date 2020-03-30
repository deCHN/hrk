package hrk

import "testing"

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
	return 0
}
