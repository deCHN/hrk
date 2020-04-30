package hrk_test

import "testing"

func TestViralAdvertising(t *testing.T) {
	tests := []struct {
		n, want int32
	}{
		{3, 9},
	}

	for _, v := range tests {
		if get := viralAdvertising(v.n); get != v.want {
			t.Errorf("Given %v, want %v, but %v.\n", v.n, v.want, get)
		}
	}

}

func viralAdvertising(n int32) int32 {
	return 0
}
