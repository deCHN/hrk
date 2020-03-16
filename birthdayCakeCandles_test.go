package hrk

import (
	"sort"
	"testing"
)

func TestBirthdayCakeCandles(t *testing.T) {
	input := []int32{3, 2, 1, 3} //Example 0
	//We have one candle of height 1, one candle of height 2, and
	//two candles of height 3. Your niece only blows out the tallest
	//candles, meaning the candles where height = 3. Because there
	//are 2 such candles, we print 2 on a new line.
	expect := int32(2)

	if get := birthdayCakeCandles(input); get != expect {
		t.Errorf("Expected %v, but %v.", expect, get)
	}

}

func birthdayCakeCandles(ar []int32) int32 {
	sort.SliceStable(ar, func(i int, j int) bool {
		return ar[i] > ar[j]
	})

	for k, v := range ar {
		if v != ar[0] {
			return int32(k)
		}
	}
	return int32(len(ar))
}
