package sort

import (
	// "fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {

	arr := []_T{34, 8, 64, 51, 32, 21}
	arr2 := []_T{8, 21, 32, 34, 51, 64}

	InsertionSort(&arr)

	for i := range arr {
		if arr[i] != arr2[i] {
			t.Errorf("Error InsertionSort")
		}
	}
}
