package sort

import (
	// "fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {

	arr := []int{34, 8, 64, 51, 32, 21}
	arr2 := []int{8, 21, 32, 34, 51, 64}

	BubbleSort(&arr)

	// 比照数组
	for i := range arr {
		if arr[i] != arr2[i] {
			t.Errorf("Error BubbleSort")
		}
	}
}
