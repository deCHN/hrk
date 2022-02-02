package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * Implement a simple text editor. The editor initially contains an empty string, . Perform  operations of the following  types:
 *
 * 1 - append - Append string  to the end of .
 * 2 - delete - Delete the last  characters of .
 * 3 - print - Print the  character of .
 * 4 - undo - Undo the last (not previously undone) operation of type  or , reverting  to the state it was in prior to that operation.
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
