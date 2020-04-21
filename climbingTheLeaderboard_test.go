package hrk_test

import (
	"fmt"
	"testing"
)

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

		if len(get) != len(v.want) {
			t.Fatalf("Length doesn't match. Want %v, but %v.", len(v.want), len(get))
		}

		for k, s := range get {
			if s != v.want[k] {
				t.Errorf("Rank doesn't match. Want %v, but %v.", v.want[k], s)
			}
		}
	}
}

func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	denseRank := func(scores []int32) []int32 {
		rank := make([]int32, len(scores)) // score indexed with rank
		i := 1                             // Rank starts from 1

		for k, v := range scores {
			scores[k] = v + 1 // To aviod 0 as score
		}

		for k, score := range scores {
			if k == 0 {
				rank[0] = score
				continue
			}
			if score == scores[k-1] {
				continue
			}
			rank[i] = score
			i++
		}

		// filter zero values away
		n := 0
		for _, x := range rank {
			if x != 0 {
				rank[n] = x
				n++
			}
		}
		rank = rank[:n]

		return rank
	}

	leaderBoard := denseRank(scores)
	fmt.Println("leader board:", leaderBoard)

	var ranking []int32
	for _, a := range alice {
		i := 0
		for r, score := range leaderBoard {
			i++
			if a > score {
				if r == 1 {
					ranking = append(ranking, 1)
				} else {
					ranking = append(ranking, int32(r-1))
				}
			}
			if i == len(leaderBoard) {
				ranking = append(ranking, int32(r+1))
			}
		}
	}

	return ranking
}
