package hrk

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

/*
 *Devendra在9号云上看到了他的教练朝他微笑。 每次教授选出Devendra单独问他一个问题，Devendra朦胧的头脑里全是他的教练和她的微笑，以至于他无法专注于其他事情。帮助他解决这个问题：
 *给你一个长度为N的列表，列表的初始值全是0。对此列表，你要进行M次查询，输出列表种最终N个值的最大值。对每次查询，给你的是3个整数——a,b和k，你要对列表中从位置a到位置b范围内的（包含a和b)的全部元素加上k。
*
* 约束条件
*
* 3 <= N <= 10^7
* 1 <= M <= 2*10^5
* 1 <= a <= b <= N
* 0 <= k <= 10^9
*/

func TestArrayManipulation(t *testing.T) {
	td := []struct {
		n       int32
		queries [][]int32
		want    int64
	}{
		/* Case 1.*/ {5, [][]int32{{1, 2, 100}, {2, 5, 100}, {3, 4, 100}}, 200},
		/* Case 2.*/ {10, [][]int32{{1, 5, 3}, {4, 8, 7}, {6, 9, 1}}, 10},
		/* Case 3.*/ {10, [][]int32{{2, 6, 8}, {3, 5, 7}, {1, 8, 1}, {5, 9, 15}}, 31},
	}

	for k, d := range td {
		get := arrayManipulation(d.n, d.queries)
		if get != d.want {
			t.Errorf("Testcase %v failed. Want: %v, but get: %v.", k, d.want, get)
		}
	}
}

/*
 * Complete the 'arrayManipulation' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */
func arrayManipulation(n int32, queries [][]int32) int64 {
	dc := make(map[int32]int64)

	for k, q := range queries {
		fmt.Printf("Query %d: %v.\n", k, q)
		for i := q[0]; i <= q[1]; i++ {
			dc[i] = dc[i] + int64(q[2])
		}
	}

	max := int64(0)
	for _, v := range dc {
		if int64(v) > max {
			max = int64(v)
		}
	}
	return max
}

func TestArrayManipulationInputs(t *testing.T) {
	inputs := []struct {
		file string
		want int64
	}{
		/* case 4 */ {"input04.txt", 7542539201},
		/* case 13 */ //{"input13.txt", 2490686975},
	}
	for _, v := range inputs {
		if get := arrayManipulationInput(v.file); get != v.want {
			t.Errorf("Test failed. Want %d, but get %d.", v.want, get)
		}
	}
}

func arrayManipulationInput(file string) int64 {
	input, err := os.Open(file)
	checkError(err)

	reader := bufio.NewReaderSize(input, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var queries [][]int32
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 3 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := arrayManipulation(n, queries)

	return result
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
