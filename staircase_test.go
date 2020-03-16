package hrk

import (
	"fmt"
	"strings"
	"testing"
)

func TestStaircase(t *testing.T) {
	staircase(6)
}

// Complete the plusMinus function below.
func staircase(n int32) {
	for i := 1; i <= int(n); i++ {
		fmt.Println(fixedLengthString(int(n), strings.Repeat("#", i)))
	}
}

func fixedLengthString(length int, str string) string {
	verb := fmt.Sprintf("%%%d.%ds", length, length)
	return fmt.Sprintf(verb, str)
}
