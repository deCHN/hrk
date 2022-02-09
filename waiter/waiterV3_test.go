package main

import (
	"testing"
)

func prmV3(n int) []int32 {
	primaries := make([]int32, 0) // [2,3,5,7,11,...]
	primaries = append(primaries, 2)

	for query := 1; query < n; query++ {
		// Find the next primary by starting from the biggest known primary
		for next := primaries[len(primaries)-1] + 1; ; next++ {
			//check if next is primary number
			if func(n int32) bool {
				for _, p := range primaries {
					if n%p == 0 {
						return false
					}
				}
				return true
			}(next) {
				// if next is a primary number , add it to the list
				primaries = append(primaries, next)
				// and start the next query
				break
			}
		}
	}

	return primaries
}

func TestPrmV3(t *testing.T) {
	for i := 1; i < 50; i++ {
		t.Logf("%d: %v", i, prmV3(i))
	}
}
