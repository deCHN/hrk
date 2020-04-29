package hrk_test

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestBeautifulDays(t *testing.T) {
	tests := []struct {
		i, j, k int32
		want    int32
	}{
		{20, 23, 6, 2}, // (20 - 2) / 6, (22 - 22 ) / 6
	}

	for _, v := range tests {
		if get := beautifulDays(v.i, v.j, v.k); get != v.want {
			t.Errorf("Give %v, want %v, but get %v.\n", v, v.want, get)
		}
	}
}

func TestStringAssign(t *testing.T) {
	str := "12340"
	var rts []byte

	size := len(str)
	for k, _ := range str {
		rts = append(rts, str[size-k-1])
	}

	fmt.Printf("Reverse \"%s\" -> \"%s\".\n", str, rts)
	if v, err := strconv.Atoi(fmt.Sprintf("%s", rts)); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}
}

func beautifulDays(i int32, j int32, k int32) int32 {

	reverse := func(x int32) int32 {
		str := strconv.Itoa(int(x))
		var rts []byte
		size := len(str)

		for k, _ := range str {
			rts = append(rts, str[size-k-1])
		}

		v, _ := strconv.Atoi(fmt.Sprintf("%s", rts))

		return int32(v)
	}

	count := 0

	for v := i; v < j; v++ {
		df := v - reverse(v)
		df = int32(math.Abs(float64(df)))

		if math.Mod(float64(df), float64(k)) == 0 {
			fmt.Printf("%v is a beautiful day.\n", v)
			count++
		} else {
			fmt.Printf("%v is not a beautiful day becasue %v/%v is not a whole number.\n", v, df, k)
		}

	}
	return int32(count)
}
