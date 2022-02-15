package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

/* Alexa has two stacks of non-negative integers, stack  and stack  where index  denotes the top of the stack. Alexa challenges Nick to play the following game:

In each move, Nick can remove one integer from the top of either stack  or stack .
Nick keeps a running sum of the integers he removes from the two stacks.
Nick is disqualified from the game if, at any point, his running sum becomes greater than some integer  given at the beginning of the game.
Nick's final score is the total number of integers he has removed from the two stacks.
Given , , and  for  games, find the maximum possible score Nick can achieve.

Example


The maximum number of values Nick can remove is . There are two sets of choices with this result.

Remove  from  with a sum of .
Remove  from  and  from  with a sum of .
Function Description
Complete the twoStacks function in the editor below.

twoStacks has the following parameters: - int maxSum: the maximum allowed sum
- int a[n]: the first stack
- int b[m]: the second stack

Returns
- int: the maximum number of selections Nick can make

Input Format

The first line contains an integer,  (the number of games). The  subsequent lines describe each game in the following format:

The first line contains three space-separated integers describing the respective values of  (the number of integers in stack ),  (the number of integers in stack ), and  (the number that the sum of the integers removed from the two stacks cannot exceed).
The second line contains  space-separated integers, the respective values of .
The third line contains  space-separated integers, the respective values of .
*/

func TestTwoStacks(t *testing.T) {
	td := []struct {
		max  int32
		a    []int32
		b    []int32
		want int32
	}{
		/* 1 */ {10, []int32{1, 2, 3, 4, 5}, []int32{6, 7, 8, 9}, 4},
		/* 2 */ {10, []int32{4, 2, 4, 6, 1}, []int32{2, 1, 8, 5}, 4},
		/* 3 */ {62, []int32{7, 15, 12, 0, 5, 18, 17, 2, 10, 15, 4, 2, 9, 15, 13, 12, 16}, []int32{12, 16, 6, 8, 16, 15, 18, 3, 11, 0, 17, 7, 6, 11, 14, 13, 15, 6, 18, 6, 16, 12, 16, 11, 16, 11}, 6},
		/* 4 */ {67, []int32{19, 9, 8, 13, 1, 7, 18, 0, 19, 19, 10, 5, 15, 19, 0, 0, 16, 12, 5, 10}, []int32{11, 17, 1, 18, 14, 12, 9, 18, 14, 3, 4, 13, 4, 12, 6, 5, 12, 16, 5, 11, 16, 8, 16, 3, 7, 8, 3, 3, 0, 1, 13, 4, 10, 7, 14}, 6},
	}

	for ti, v := range td {
		get := twoStacks(v.max, v.a, v.b)

		if get != v.want {
			t.Errorf("Case %d ✗\t want %v but get %v.", ti+1, v.want, get)
		} else {
			t.Logf("Case %d ✔", ti+1)
		}
	}
}

/*
 * Complete the 'twoStacks' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER maxSum
 *  2. INTEGER_ARRAY a
 *  3. INTEGER_ARRAY b
 */
func twoStacks(maxSum int32, a []int32, b []int32) int32 {
	la := len(a)
	lb := len(b)

	ai := 0
	bi := 0

	sum := int32(0)
	var n int32

	for n = 0; sum < maxSum; n++ {
		if ai < la && bi < lb {
			if a[ai] < b[bi] {
				sum += a[ai]
				ai++
			} else {
				sum += b[bi]
				bi++
			}
		} else if ai < la {
			sum += a[ai]
			ai++
		} else if bi < lb {
			sum += b[bi]
			bi++
		} else {
			break
		}
	}

	if sum > maxSum {
		return n - 1
	}

	return n
}

func TestTwoStacksInputs(t *testing.T) {
	input, err := os.Open("./input01.txt")
	checkError(err)

	reader := bufio.NewReader(input)
	gTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	g := int32(gTemp)

	want, err := os.Open("./want01.txt")
	checkError(err)

	scanner := bufio.NewScanner(want)

	for gItr := 0; gItr < int(g); gItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		maxSumTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
		checkError(err)
		maxSum := int32(maxSumTemp)

		aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var a []int32

		for i := 0; i < int(n); i++ {
			aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
			checkError(err)
			aItem := int32(aItemTemp)
			a = append(a, aItem)
		}

		bTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var b []int32

		for i := 0; i < int(m); i++ {
			bItemTemp, err := strconv.ParseInt(bTemp[i], 10, 64)
			checkError(err)
			bItem := int32(bItemTemp)
			b = append(b, bItem)
		}

		result := twoStacks(maxSum, a, b)
		if !scanner.Scan() {
			t.Fatal("Indexes of inputs differ from wants.")
		}

		get := strconv.Itoa(int(result))

		if get != scanner.Text() {
			t.Errorf("Case %d ✗\twant %#v but get %#v.", gItr+1, scanner.Text(), get)
		}

		t.Logf("Case %d ✔", gItr+1)
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
