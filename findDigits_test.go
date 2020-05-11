package hrk_test

import (
	"fmt"
	"strconv"
	"testing"
)

func TestFindDigits(t *testing.T) {
	tests := []struct {
		n, want int32
	}{
		{12, 2},
		{1012, 3},
	}

	for _, v := range tests {
		if get := findDigits(v.n); get != v.want {
			t.Errorf("Given %v, want %v, but get %v.\n", v.n, v.want, get)
		}
	}
}

func findDigits(n int32) int32 {
	ns := strconv.Itoa(int(n))
	count := int32(0)

	for _, v := range ns {
		i, err := strconv.Atoi(string(v))

		if err != nil {
			return -1
		}

		if i != 0 && n%int32(i) == 0 {
			count++
		}
	}

	fmt.Println(count)

	return count
}
