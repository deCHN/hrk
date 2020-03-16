package hrk

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestTimeConversion(t *testing.T) {
	tt := [][2]string{
		[2]string{"07:05:45PM", "19:05:45"},
		[2]string{"07:05:45AM", "07:05:45"},
		[2]string{"00:05:45PM", "12:05:45"},
		[2]string{"00:05:45AM", "00:05:45"},
		[2]string{"00:00:00AM", "00:00:00"},
		[2]string{"00:00:00PM", "12:00:00"},
		[2]string{"12:00:00AM", "00:00:00"},
		[2]string{"12:00:00PM", "12:00:00"},
		[2]string{"12:45:54PM", "12:45:54"}, // Is this a valid time?
	}

	for _, tc := range tt {
		if get := timeConversion(tc[0]); get != tc[1] {
			t.Errorf("Expect %v, but %v.", tc[1], get)
		}
	}
}

func timeConversion(s string) string {
	const (
		pm = "PM"
		am = "AM"
		sp = ":"
	)
	hhs := strings.SplitN(s, sp, 2)

	var hh int
	if v, err := strconv.Atoi(hhs[0]); err != nil {
		return ""
	} else if strings.Contains(s, pm) {
		if v == 12 {
			hh = v
		} else {
			hh = (v + 12) % 24
		}
	} else {
		hh = v % 12
	}
	hhs[0] = fmt.Sprintf("%02d", hh)

	return strings.TrimRight(strings.Join(hhs, sp), am+pm)
}
