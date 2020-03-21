package hrk

import "testing"

func TestGetTotalX(t *testing.T) {

	a := []int32{2, 4}
	b := []int32{16, 32, 96}

	expect := int32(3)

	if get := getTotalX(a, b); get != expect {
		t.Fail()
	}
}

func getTotalX(a []int32, b []int32) int32 {
	//x % a[i] -> 0
	//b[i] % x -> 0

	return 0
}

func isFactorOf(f int32, fc []int32) bool {
	for _, v := range fc {
		if v%f != 0 {
			return false
		}
	}

	return true
}
