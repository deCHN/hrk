package hrk_test

import (
	"bufio"
	"container/ring"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestSaveThePrisoner(t *testing.T) {
	tests := []struct {
		n, m, s int32
		want    int32
	}{
		{7, 19, 2, 6},
		{3, 7, 3, 3},
	}

	for _, v := range tests {
		if get := saveThePrisoner(v.n, v.m, v.s); get != v.want {
			t.Errorf("Given %v, want %v, but %v.\n", v, v.want, get)
		}
	}
}

func TestSaveThePrisonerInput00(ti *testing.T) {
	f, _ := os.Open("./input00.txt")
	reader := bufio.NewReaderSize(f, 1024*1024)

	stdout, err := os.Create("./output00.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nms := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nms[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(nms[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		sTemp, err := strconv.ParseInt(nms[2], 10, 64)
		checkError(err)
		s := int32(sTemp)

		result := saveThePrisoner(n, m, s)

		fmt.Fprintf(writer, "%d\n", result)
	}

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

/*
 * https://www.hackerrank.com/challenges/save-the-prisoner/problem
 *
 * n: the number of prisoners. 1 <=n <= 10^9
 * m: the number of sweets. 1 <=m <= 10^9
 * s: the chair number to start passing out treats at. 1 <= s <= n
 */
func saveThePrisoner(n int32, m int32, s int32) int32 {
	r := ring.New(int(n))

	for i := int32(1); i <= n; i++ {
		r.Value = i
		r = r.Next()
	}

	for i := 0; i < r.Len(); i++ {
		if r.Value == s {
			break
		}
		r = r.Next()
	}

	r = r.Move(int(m - 1))

	return r.Value.(int32)
}
