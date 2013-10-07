package klib

import (
	"fmt"
	"testing"
)

func TestKVector(t *testing.T) {

	// Test Pointer type
	var vec *KVector = NewKVector() // size 1

	// Test Insert
	for i := 0; i < 20; i++ {
		vec.PushBack(i)
		if vec.Size() != i+1 {
			//t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
			t.Errorf("vec.Size() != %d, RealSize:%d", i+1, vec.Size())
		}
	}

	// Test Index value
	var tmp int = vec.Get(0).(int)
	if tmp != 0 {
		t.Errorf("Index 0 Value != 0,  Is %d", tmp)
	}

	// Test Remove
	for i := 0; i < 20; i++ {
		vec.PopBack()
		if vec.Size() != 20-1-i {
			//t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
			t.Errorf("vec.Size() != %d, RealSize:%d", 20-1-i, vec.Size())
		}
		fmt.Printf("Now The Size : %d,  Now The Capacity: %d \n", vec.Size(), vec.Capacity())
	}

}
