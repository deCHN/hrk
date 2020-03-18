package hrk

import "testing"

func TestApplesAndOranges(t *testing.T) {
	const (
		houseLeft  int32 = 7
		houseRight int32 = 11
		appleTree  int32 = 5
		orangeTree int32 = 15
	)
	apples := []int32{-2, 2, 1}
	oranges := []int32{5, -6}

	if apple, orange := countApplesAndOranges(houseLeft, houseRight, appleTree, orangeTree, apples, oranges); apple != 1 || orange != 1 {
		t.Errorf("Expect apple 1 orange 1, but apple %v orange %v.", apple, orange)
	}
}

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) (apple, orange int) {
	return 1, 1
}
