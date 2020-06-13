package hrk_test

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestQueensAttack(t *testing.T) {
	tests := []struct {
		n, k      int32
		r_q, c_q  int32
		obstacles [][]int32
		want      int32
	}{
		{4, 0, 4, 4, [][]int32{nil}, 9},
		{5, 3, 4, 3, [][]int32{[]int32{5, 5}, []int32{4, 2}, []int32{2, 3}}, 10},
	}

	for _, v := range tests {
		if get := queensAttack(v.n, v.k, v.r_q, v.c_q, v.obstacles); get != v.want {
			t.Errorf("Given %v, want %v, but get %v.\n", v.n, v.want, get)
		}
	}
}

// https://www.hackerrank.com/challenges/queens-attack-2/problem
func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	// row - column pair
	type point [2]int32

	type direction func(p point) point

	sw := func(p point) point { return point{p[0] - 1, p[1] - 1} }
	south := func(p point) point { return point{p[0] - 1, p[1]} }
	se := func(p point) point { return point{p[0] - 1, p[1] + 1} }
	east := func(p point) point { return point{p[0], p[1] + 1} }
	ne := func(p point) point { return point{p[0] + 1, p[1] + 1} }
	north := func(p point) point { return point{p[0] + 1, p[1]} }
	nw := func(p point) point { return point{p[0] + 1, p[1] - 1} }
	west := func(p point) point { return point{p[0], p[1] - 1} }

	// check checks if the point p is a valid location
	check := func(p point) bool {
		// within the board
		if p[0] < 1 || p[1] < 1 {
			return false
		}

		if p[0] > n || p[1] > n {
			return false
		}

		// not on obstacles
		for _, ob := range obstacles {
			if ob == nil {
				break
			}

			if p[0] == ob[0] && p[1] == ob[1] {
				return false
			}
		}

		return true
	}

	var directions []direction = []direction{sw, south, se, east, ne, north, nw, west}

	moves := int32(0)

	for _, dir := range directions {
		p := point{r_q, c_q}
		for {
			p = dir(p)
			if check(p) {
				moves++
				continue
			}
			break
		}
	}

	return moves
}

func TestQueenAttackInput(t *testing.T) {
	input, err := os.Open("./input13.txt")
	checkError(err)
	reader := bufio.NewReaderSize(input, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	r_qC_q := strings.Split(readLine(reader), " ")

	r_qTemp, err := strconv.ParseInt(r_qC_q[0], 10, 64)
	checkError(err)
	r_q := int32(r_qTemp)

	c_qTemp, err := strconv.ParseInt(r_qC_q[1], 10, 64)
	checkError(err)
	c_q := int32(c_qTemp)

	var obstacles [][]int32
	for i := 0; i < int(k); i++ {
		obstaclesRowTemp := strings.Split(readLine(reader), " ")

		var obstaclesRow []int32
		for _, obstaclesRowItem := range obstaclesRowTemp {
			obstaclesItemTemp, err := strconv.ParseInt(obstaclesRowItem, 10, 64)
			checkError(err)
			obstaclesItem := int32(obstaclesItemTemp)
			obstaclesRow = append(obstaclesRow, obstaclesItem)
		}

		if len(obstaclesRow) != int(2) {
			panic("Bad input")
		}

		obstacles = append(obstacles, obstaclesRow)
	}

	get := queensAttack(n, k, r_q, c_q, obstacles)

	f, err := os.Open("./want13.txt")
	checkError(err)
	defer f.Close()

	r, err := bufio.NewReader(f).ReadString('\n')
	if err != io.EOF {
		checkError(err)
	}
	want, err := strconv.Atoi(r)
	checkError(err)

	if get != int32(want) {
		t.Errorf("Given input%v.txt, want %v, but get %v.\n", "13", want, get)
	}
}

func BenchmarkQueensAttack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queensAttack(5, 3, 4, 3, [][]int32{[]int32{5, 5}, []int32{4, 2}, []int32{2, 3}})
	}
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
