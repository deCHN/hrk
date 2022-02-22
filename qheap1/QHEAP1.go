package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	line := ""
	undo := make([]string, 0)

	for scanner.Scan() {
		op := strings.Split(scanner.Text(), " ")

		switch op[0] {
		case "1":
			undo = append(undo, line)
			line += op[1]
		case "2":
			undo = append(undo, line)
			i, _ := strconv.Atoi(op[1])
			line = line[:len(line)-i]
		case "3":
			i, _ := strconv.Atoi(op[1])
			fmt.Println(string(line[i-1]))
		case "4":
			line = undo[len(undo)-1]
			undo = undo[:(len(undo) - 1)]
		default:
			continue
		}
	}
}
