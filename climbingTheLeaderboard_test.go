package hrk_test

import "testing"

func TestClimbingLeaderBoard(t *testing.T) {
	tests := []struct {
		scores, alice []int32
		want          []int32
	}{
		{[]int32{100, 100, 50, 40, 40, 20, 10}, []int32{5, 25, 50, 120}, []int32{6, 4, 2, 1}},
		{[]int32{100, 90, 90, 80, 75, 60}, []int32{50, 65, 77, 90, 102}, []int32{6, 5, 4, 2, 1}},
	}

	for _, v := range tests {
		get := climbingLeaderboard(v.scores, v.alice)
		if get == nil {
			t.Error("Get nil.")
		}

		for k, s := range get {
			if s != v.want[k] {
				t.Errorf("Want %v, but %v.", v.want[k], s)
			}
		}
	}
}

func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	denseRank := func(scores []int32) map[int32]int {
		rank := make(map[int32]int) // score to rank map

		i := 1 // Rank starts from 1
		for _, score := range scores {
			if _, ok := rank[score]; ok {
				continue
			}
			rank[score] = i
			i++
		}
		return rank
	}

	rank := denseRank(scores)

	var ranking []int32
	for _, a := range alice {
		i := 0
		for score, r := range rank {
			i++
			if a > score {
				if r == 1 {
					ranking = append(ranking, 1)
				} else {
					ranking = append(ranking, int32(r-1))
				}
			}
			if i == len(rank) {
				ranking = append(ranking, int32(r+1))
			}
		}
	}

	return ranking
}
