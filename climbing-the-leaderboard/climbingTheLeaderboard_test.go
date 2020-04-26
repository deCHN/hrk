package hrk_test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
				t.Errorf("Rank doesn't match. Alice's score is %v, want %v, but %v.", v.alice[k], v.want[k], s)
			}
		}
	}
}

// Got runtime error when testing online.
// But local run's output matches 100% of wanted result.
func TestClimbingLeaderBoardInput6(t *testing.T) {
	input, err := os.Open("./input06.txt")
	if err != nil {
		t.Error(err)
	}
	defer input.Close()

	reader := bufio.NewReaderSize(input, 2*1024*1024)
	fmt.Printf("Reader's buffer size = %v byte(s)\n", reader.Size())

	stdout, err := os.Create("./output06.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	scoresCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	scoresTemp := strings.Split(readLine(reader), " ")

	var scores []int32

	for i := 0; i < int(scoresCount); i++ {
		scoresItemTemp, err := strconv.ParseInt(scoresTemp[i], 10, 64)
		checkError(err)
		scoresItem := int32(scoresItemTemp)
		scores = append(scores, scoresItem)
	}

	aliceCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	aliceTemp := strings.Split(readLine(reader), " ")

	var alice []int32

	for i := 0; i < int(aliceCount); i++ {
		aliceItemTemp, err := strconv.ParseInt(aliceTemp[i], 10, 64)
		checkError(err)
		aliceItem := int32(aliceItemTemp)
		alice = append(alice, aliceItem)
	}

	result := climbingLeaderboard(scores, alice)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// dense rank is a sort list with the index mapped to the rank
// from high(best) to low.
type denseRank []int32

func (dr denseRank) getRanking(score int32) int32 {

	rank := int32(-1)
	for r, p := range dr {
		if score >= p {
			rank = int32(r)
			break
		}
	}

	if rank == -1 {
		return int32(len(dr) + 1)
	}

	return rank + 1 //r is indexed from 0
}

//https://www.hackerrank.com/challenges/climbing-the-leaderboard/problem
func climbingLeaderboard(scores []int32, alice []int32) []int32 {

	dr := func(scores []int32) denseRank {
		rank := make([]int32, len(scores)) // score indexed with rank

		i := 1 // Rank starts from 1

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

		return rank
	}

	leaderBoard := dr(scores)

	var ranking []int32
	for _, score := range alice {
		ranking = append(ranking, leaderBoard.getRanking(score))
	}

	return ranking
}
