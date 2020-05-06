package hrk_test

import (
	"bufio"
	"container/ring"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func TestSaveThePrisonerTD(t *testing.T) {
	tests := []struct {
		n, m, s int32
		want    int32
	}{
		//{7, 19, 2, 6},
		//{3, 7, 3, 3},
		//{352926151, 380324688, 94730870, 122129406},
		//{208526924, 756265725, 150817879, 72975907},
		{499999999, 999999998, 2, 499999999}, //1
		//{4, 8, 2, 1},
		//{499999999, 999999998, 2, 1},         //1
		//{999999999, 999999999, 1, 999999999}, //0
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

func saveThePrisoner(n int32, m int32, s int32) int32 {
	m = m % n

	if m+s <= n {
		return s + m - 1
	}

	if (m+s)%n-1 == 0 {
		return s
	}

	return (m+s)%n - 1
}

/*
 * https://www.hackerrank.com/challenges/save-the-prisoner/problem
 *
 * n: the number of prisoners. 1 <=n <= 10^9
 * m: the number of sweets. 1 <=m <= 10^9
 * s: the chair number to start passing out treats at. 1 <= s <= n
 */
func saveThePrisonerWithRing(n int32, m int32, s int32) int32 {
	r := ring.New(int(n))
	PrintMemUsage()

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

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
