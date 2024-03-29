package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func init() {
	log.SetOutput(ioutil.Discard)
}

/*
 * https://www.hackerrank.com/challenges/waiter/problem?utm_campaign=challenge-recommendation&utm_medium=email&utm_source=24-hour-campaign
 *
 * You are a waiter at a party. There is a pile of numbered plates. Create an empty  array. At each iteration, , remove each plate from the top of the stack in order. Determine if the number on the plate is evenly divisible by the  prime number. If it is, stack it in pile . Otherwise, stack it in stack . Store the values in  from top to bottom in . In the next iteration, do the same with the values in stack . Once the required number of iterations is complete, store the remaining values in  in , again from top to bottom. Return the  array.
 *
 *Example
 *
 *
 *
 *An abbreviated list of primes is . Stack the plates in reverse order.
 *
 *
 *
 *Begin iterations. On the first iteration, check if items are divisible by .
 *
 *
 *Move  elements to .
 *
 *
 *On the second iteration, test if  elements are divisible by .
 *
 *
 *Move  elmements to .
 *
 *
 *And on the third iteration, test if  elements are divisible by .
 *
 *
 *Move  elmements to .
 *
 *
 *All iterations are complete, so move the remaining elements in , from top to bottom, to .
 *
 *. Return this list.
 *
 *Function Description
 *
 *Complete the waiter function in the editor below.
 *
 *waiter has the following parameters:
 *
 *int number[n]: the numbers on the plates
 *int q: the number of iterations
 *Returns
 *
 *int[n]: the numbers on the plates after processing
 *Input Format
 *
 *The first line contains two space separated integers,  and .
 *The next line contains  space separated integers representing the initial pile of plates, i.e., .
 */
func TestWaiter(t *testing.T) {
	td := []struct {
		number []int32
		q      int32
		want   []int32
	}{
		/* 1 */ {[]int32{4, 4, 9, 3, 3}, 2, []int32{4, 4, 3, 3, 9}},
		/* 2 */ {[]int32{3, 4, 7, 6, 5}, 1, []int32{4, 6, 3, 7, 5}},
		/* 3 */ {[]int32{2, 3, 4, 5, 6, 7}, 3, []int32{2, 4, 6, 3, 5, 7}},
		/* 4 */ {[]int32{3, 3, 4, 4, 9}, 2, []int32{4, 4, 9, 3, 3}},
	}

	for k, tc := range td {
		get := waiter(tc.number, tc.q)
		log.Println("GET:", get)

		if !reflect.DeepEqual(tc.want, get) {
			t.Errorf("Case %d failed. Want %v, but get %v.", k, tc.want, get)
		}
	}
}

/*
 * Complete the 'waiter' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY number
 *  2. INTEGER q
 */
func waiter(number []int32, q int32) []int32 {
	answer := make([]int32, 0)
	prm := prm(int(q)) // 32bit -> 32bit or 64bit depends on maschin

	a := number

	for i := 0; i < int(q); i++ {
		number = a
		a = make([]int32, 0)

		for _, v := range number {
			if v%prm[i] == 0 {
				answer = append(answer, v)
			} else {
				a = append([]int32{v}, a...)
			}
		}
	}

	for i := len(a) - 1; i >= 0; i-- {
		answer = append(answer, a[i])
	}
	return answer
}

// func prm returns the first n prim numbers in a list
func prm(n int) []int32 {
	primes := make([]int32, 0)
	primes = append(primes, 2, 3)

	for query := len(primes); query < n; query++ {
		// Find the next prim by starting from the biggest known prim
		for next := primes[len(primes)-1] + 2; ; next++ {
			//check if next is prim number
			if func(n int32) bool {
				for _, p := range primes {
					if n%p == 0 {
						return false
					}
					if p*p > n {
						return true
					}
				}
				return true
			}(next) {
				// if next is a prim number , add it to the list
				primes = append(primes, next)
				// and start the next query
				break
			}
		}
	}

	return primes // [2,3,5,7,11,...]
}

func TestPrm(t *testing.T) {
	primaries := prm(30000)

	for _, v := range primaries {
		if !big.NewInt(int64(v)).ProbablyPrime(0) {
			t.Errorf("%d is probably not a prim number.", v)
		}
	}
}

func BenchmarkPrm1K(b *testing.B) {
	benchmarkPrims(1000, b)
}

func BenchmarkPrm10k(b *testing.B) {
	benchmarkPrims(10*1000, b)
}

func BenchmarkPrm100k(b *testing.B) {
	benchmarkPrims(100*1000, b)
}

func benchmarkPrims(x int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		prm(x)
	}
}

func TestWaiterInputs(t *testing.T) {
	input, err := os.Open("./input12.txt")
	checkError(err)

	reader := bufio.NewReaderSize(input, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	qTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	q := int32(qTemp)

	numberTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var number []int32

	for i := 0; i < int(n); i++ {
		numberItemTemp, err := strconv.ParseInt(numberTemp[i], 10, 64)
		checkError(err)
		numberItem := int32(numberItemTemp)
		number = append(number, numberItem)
	}

	result := waiter(number, q)

	want, err := os.Open("./want12.txt")
	checkError(err)

	wscanner := bufio.NewScanner(want)
	wants := make([]int32, 0)

	for wscanner.Scan() {
		wantItem, err := strconv.Atoi(wscanner.Text())
		checkError(err)

		wants = append(wants, int32(wantItem))
	}

	for i, resultItem := range result {
		if resultItem != wants[i] {
			t.Errorf("Test failed. Line %d diff, want %v, but get %v.\n", i, wants[i], resultItem)
		}
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
