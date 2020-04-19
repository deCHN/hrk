package hrk_test

import "testing"

func TestClimbingLeaderBoard(t *testing.T) {
	tests := []struct {
		scores, alice []int32
		want          []int32
	}{
		{[]int32{100, 100, 50, 40, 40, 20, 10}, []int32{5, 25, 50, 120}, []int32{5, 25, 50, 120}},
		{[]int32{100, 90, 90, 80, 75, 60}, []int32{50, 65, 77, 90, 102}, []int32{6, 5, 4, 2, 1}},
	}

	for _, v := range tests {
		get := climbingLeaderboard(v.scores, v.alice)
		if get == nil {
			t.Error("Get nil.")
		}

		for k, s := range get {
			if s != v.want[k] {
				t.Errorf("Want %v, but %v.", v.want, s)
			}
		}
	}
}

func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	denseRank := func(scores []int32) map[int]int32 {
		rank := make(map[int]int32) // rank-value

		//for k, v := range scores {
		//if rank[v]

		//}
		return nil
	}

	rank := denseRank(scores)

	for k, v := range alice {

	}

	return nil

}
