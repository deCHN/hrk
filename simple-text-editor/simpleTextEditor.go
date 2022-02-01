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

	for scanner.Scan() {
		op := strings.Split(scanner.Text(), " ")

		switch op[0] {
		case "1":
			line += op[1]
			fmt.Println("append", line)
		case "2":
			i, _ := strconv.Atoi(op[1])
			line = line[:len(line)-i]
			fmt.Println("delete", line)
		case "3":
			i, _ := strconv.Atoi(op[1])
			fmt.Println("print", string(line[i-1]))
		case "4":
			fmt.Println("undo", line)
		default:
			continue
		}
	}
}
