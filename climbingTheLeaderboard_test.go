package hrk_test

import (
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

// dense rank is a sort list with the index mapped to the rank
// from high(best) to low.
type denseRank []int32

func (dr denseRank) getRanking(score int32) int32 {
	rank := int32(-1)
	for r, s := range dr {
		if score > s {
			if r == 0 {
				rank = 1
			} else {
				rank = int32(r)
			}
		}
	}

	if rank == -1 {
		rank = int32(len(dr) + 1)
	}
	return rank
}

func climbingLeaderboard(scores []int32, alice []int32) []int32 {

	dr := func(scores []int32) denseRank {
		rank := make([]int32, len(scores)) // score indexed with rank

		i := 1 // Rank starts from 1

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
				rank[n] = x - 1 // Set back to original value
				n++
			}
		}
		rank = rank[:n]

		return rank
	}

	leaderBoard := dr(scores)

	var ranking []int32
	for _, score := range alice {
		ranking = append(ranking, leaderBoard.getRanking(score))
	}

	return ranking
}
