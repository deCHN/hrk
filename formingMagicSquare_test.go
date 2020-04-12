package hrk_test

import "testing"

func TestFormingMagicSquare(t *testing.T) {

	tests := []struct {
		s    [][]int32
		want int32
	}{
		{[][]int32{
			{5, 3, 4}, //{8, 3, 4}, -> |5 - 8| = 3
			{1, 5, 8}, //{1, 5, 9}, -> |8 - 9| = 1
			{6, 4, 2}, //{6, 7 ,2}, -> |4 - 7| = 3
		}, 7}, // 3 + 1 + 3 = 7
		{[][]int32{
			{4, 9, 2}, //{4, 9 ,2},
			{3, 5, 7}, //{3, 5, 7},
			{8, 1, 5}, //{8, 1, 6}, -> |5 - 6| = 1
		}, 1}, // 0 + 0 + 1 = 1
	}
}

func formingMagicSquare(s [][]int32) int32 {

}
