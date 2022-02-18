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
 * https://www.hackerrank.com/challenges/truck-tour/problem?utm_campaign=challenge-recommendation&utm_medium=email&utm_source=24-hour-campaign
 *
 * Suppose there is a circle. There are  petrol pumps on that circle. Petrol pumps are numbered  to  (both inclusive). You have two pieces of information corresponding to each of the petrol pump: (1) the amount of petrol that particular petrol pump will give, and (2) the distance from that petrol pump to the next petrol pump.
 *
 * Initially, you have a tank of infinite capacity carrying no petrol. You can start the tour at any of the petrol pumps. Calculate the first point from where the truck will be able to complete the circle. Consider that the truck will stop at each of the petrol pumps. The truck will move one kilometer for each litre of the petrol.
 *
 * Input Format
 *
 * The first line will contain the value of .
 * The next  lines will contain a pair of integers each, i.e. the amount of petrol that petrol pump will give and the distance between that petrol pump and the next petrol pump.
 */

func TestTruckTour(t *testing.T) {
	td := []struct {
		petrolpumps [][]int32
		want        int32
	}{
		{[][]int32{[]int32{1, 5}, []int32{10, 3}, []int32{3, 4}}, 1},
	}

	for i, tc := range td {
		get := truckTour(tc.petrolpumps)

		if get != tc.want {
			t.Errorf("Case %d failed. Want %d, but get %d.", i+1, tc.want, get)
		}
	}
}

/*
 * Complete the 'truckTour' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY petrolpumps as parameter.
 */
func truckTour(petrolpumps [][]int32) int32 {
	return 0
}

func TestTruckTourInputs(t *testing.T) {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var petrolpumps [][]int32
	for i := 0; i < int(n); i++ {
		petrolpumpsRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var petrolpumpsRow []int32
		for _, petrolpumpsRowItem := range petrolpumpsRowTemp {
			petrolpumpsItemTemp, err := strconv.ParseInt(petrolpumpsRowItem, 10, 64)
			checkError(err)
			petrolpumpsItem := int32(petrolpumpsItemTemp)
			petrolpumpsRow = append(petrolpumpsRow, petrolpumpsItem)
		}

		if len(petrolpumpsRow) != 2 {
			panic("Bad input")
		}

		petrolpumps = append(petrolpumps, petrolpumpsRow)
	}

	result := truckTour(petrolpumps)

	fmt.Fprintf(writer, "%d\n", result)

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
