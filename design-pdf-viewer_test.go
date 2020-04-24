package hrk_test

import (
	"testing"
)

func TestDesignerPdfViewer(t *testing.T) {
	tests := []struct {
		h    []int32
		word string
		want int32
	}{
		{[]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}, "abc", 9},
		{[]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}, "zaba", 28},
	}

	for _, v := range tests {
		if get := designerPdfViewer(v.h, v.word); get != v.want {
			t.Errorf("Given %v, want %v, but get %v.\n", v.word, v.want, get)
		}
	}
}

// https://www.hackerrank.com/challenges/designer-pdf-viewer/problem
func designerPdfViewer(h []int32, word string) int32 {
	var max int32

	for _, c := range word {
		if v := h[c-97]; v > max {
			max = v
		}
	}
	return max * int32(len(word))
}
