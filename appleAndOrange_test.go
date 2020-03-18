package hrk

import "testing"

func TestCountApplesAndOranges(t *testing.T) {

	if apple, orange := countApplesAndOranges(7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6}); apple != 1 || orange != 1 {
		t.Errorf("Expect apple 1 orange 1, but apple %v orange %v.", apple, orange)
	}
}

//takes (houseLeft, houseRight, appleTree, orangeTree, applesDistance, orangesDistance)
func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) (int, int) {
	var apple, orange int

	for _, da := range apples {
		if a+da >= s && a+da <= t {
			apple++
		}
	}

	for _, db := range oranges {
		if b+db >= s && b+db <= t {
			orange++
		}
	}

	return apple, orange
}
