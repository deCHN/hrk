package hrk

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

/*
 *A bracket is considered to be any one of the following characters: (, ), {, }, [, or ].
 *
 *Two brackets are considered to be a matched pair if the an opening bracket (i.e., (, [, or {) occurs to the left of a closing bracket (i.e., ), ], or }) of the exact same type. There are three types of matched pairs of brackets: [], {}, and ().
 *
 *A matching pair of brackets is not balanced if the set of brackets it encloses are not matched. For example, {[(])} is not balanced because the contents in between { and } are not balanced. The pair of square brackets encloses a single, unbalanced opening bracket, (, and the pair of parentheses encloses a single, unbalanced closing square bracket, ].
 *
 *By this logic, we say a sequence of brackets is balanced if the following conditions are met:
 *
 *It contains no unmatched brackets.
 *The subset of brackets enclosed within the confines of a matched pair of brackets is also a matched pair of brackets.
 *Given  strings of brackets, determine whether each sequence of brackets is balanced. If a string is balanced, return YES. Otherwise, return NO.
 *
 *Function Description
 *
 *Complete the function isBalanced in the editor below.
 *
 *isBalanced has the following parameter(s):
 *
 *string s: a string of brackets
 *Returns
 *
 *string: either YES or NO
 *Input Format
 *
 *The first line contains a single integer , the number of strings.
 *Each of the next  lines contains a single string , a sequence of brackets.
 */
func TestIsBalanced(t *testing.T) {
	td := []struct {
		n    uint
		in   string
		want string
	}{
		{1, "{[()]}", "YES"},
		{2, "{[(])}", "NO"},
		{3, "{{{[[(())]]}}}", "YES"},
	}

	for _, v := range td {
		get := isBalanced(v.in)
		if get != v.want {
			t.Errorf("Testcase %d failed. Want %v, but get %v.\n", v.n, v.want, get)
		}
	}
}

/*
 * Complete the 'isBalanced' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func isBalanced(s string) string {
	balanced := make([]rune, 0)

	for _, c := range s {
		switch c {
		case '(', '[', '{':
			balanced = append(balanced, c)
		case ')', ']', '}':
			if len(balanced) == 0 {
				return "NO"
			}

			if func(p rune, q rune) bool {
				if p == '(' && q == ')' {
					return true
				} else if p == '[' && q == ']' {
					return true
				} else if p == '{' && q == '}' {
					return true
				}
				return false
			}(balanced[len(balanced)-1], c) {
				balanced = balanced[:len(balanced)-1]
			} else {
				return "NO"
			}
		}
	}

	return "YES"
}

func TestIsBalancedInputs(t *testing.T) {
	file, err := os.Open("./input04.txt")
	checkError(err)

	want, err := os.Open("./want04.txt")
	checkError(err)

	reader := bufio.NewReaderSize(file, 16*1024*1024)
	wantr := bufio.NewReader(want)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	ti32 := int32(tTemp)

	for tItr := 0; tItr < int(ti32); tItr++ {
		s := readLine(reader)
		str, err := wantr.ReadBytes('\n')
		checkError(err)

		if string(str) != isBalanced(s) {
			t.Errorf("Case %d failed.", tItr)
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
