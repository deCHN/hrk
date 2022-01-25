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

func TestDynamicArray(t *testing.T) {
	td := []struct {
		n       int32
		queries [][]int32
		want    []int32
	}{
		{2, [][]int32{{1, 0, 5}, {1, 1, 7}, {1, 0, 3}, {2, 1, 0}, {2, 1, 1}}, []int32{7, 3}},
	}

	for k, tc := range td {
		get := dynamicArray(tc.n, tc.queries)

		if len(get) != len(tc.want) {
			t.Errorf("Test %d failed. Expected results size is %d, but get %d.", k, tc.want, len(get))
		}
		for i, v := range get {
			if v != tc.want[i] {
				t.Errorf("Test %d failed at %d-th element. Exepect array %v, but get %v.", k, i, tc.want, get)
			}
		}
	}
}

func TestDynamicArrayInput() {
	file, err := os.Open("./input00.txt")
	checkError(err)
	reader := bufio.NewReaderSize(file, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	qTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	q := int32(qTemp)

	var queries [][]int32
	for i := 0; i < int(q); i++ {
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

	result := dynamicArray(n, queries)

	for i, resultItem := range result {
		fmt.Printf("%d", resultItem)

		if i != len(result)-1 {
			fmt.Printf("\n")
		}
	}

	fmt.Printf("\n")
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
 * https://www.hackerrank.com/challenges/dynamic-array/problem?utm_campaign=challenge-recommendation&utm_medium=email&utm_source=24-hour-campaign
 * Declare a 2-dimensional array, , of  empty arrays. All arrays are zero indexed.
 * Declare an integer, , and initialize it to .
 *
 * There are  types of queries, given as an array of strings for you to parse:
 *
 * Query: 1 x y
 * Let .
 * Append the integer  to .
 * Query: 2 x y
 * Let .
 * Assign the value  to .
 * Store the new value of  to an answers array.
 * Note:  is the bitwise XOR operation, which corresponds to the ^ operator in most languages. Learn more about it on Wikipedia.  is the modulo operator.
 * Finally, size(arr[idx]) is the number of elements in arr[idx]
 *
 * Function Description
 *
 * Complete the dynamicArray function below.
 *
 * dynamicArray has the following parameters:
 * - int n: the number of empty arrays to initialize in
 * - string queries[q]: query strings that contain 3 space-separated integers
 *
 * Returns
 *
 * int[]: the results of each type 2 query in the order they are presented
 * Input Format
 *
 * The first line contains two space-separated integers, , the size of  to create, and , the number of queries, respectively.
 * Each of the  subsequent lines contains a query string,
 */
func dynamicArray(n int32, queries [][]int32) []int32 {
	arr := make([][]int32, n)
	r := make([]int32, 0)
	lastAnswer := int32(0)
	idx := int32(0)

	for _, q := range queries {
		idx = q[1] ^ lastAnswer%n
		if q[0] == int32(1) {
			arr[idx] = append(arr[idx], q[2])
		} else if q[0] == 2 {
			lastAnswer = arr[idx][q[2]%int32(len(arr[idx]))]
			r = append(r, lastAnswer)
		} else {
			return nil
		}
	}

	return r
}
