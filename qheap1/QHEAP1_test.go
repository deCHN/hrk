package hrk

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

/*
 * https://www.hackerrank.com/challenges/qheap1/problem
 *
 * This question is designed to help you get a better understanding of basic heap operations.
 *
 * There are 3 types of query:
 *
 * "1 v " - Add an element v to the heap.
 * "2 v " - Delete the element v from the heap.
 * "3" - Print the minimum of all the elements in the heap.
 *
 * NOTE: It is guaranteed that the element to be deleted will be there in the heap. Also, at any instant, only distinct elements will be in the heap.
 *
 * Input Format
 *
 * The first line contains the number of queries, Q.
 * Each of the next Q lines contains one of the 3 types of query.
 */

func TestQHEAP1(t *testing.T) {
	rin := bufio.NewScanner(os.Stdin)

	for rin.Scan() {
		fmt.Println(rin.Text())
	}
}
