package klib

import (
	// "fmt"
	"testing"
)

// func TestChannel(t *testing.T) {
// 	// defer func() {
// 	// 	if err := recover(); err != nil {
// 	// 		fmt.Printf("Work failed: ", err)
// 	// 	}
// 	// }()
// 	// panic(1 == 2)

// 	var tc chan int = make(chan int, 20)
// 	for i := 0; i < 20; i++ {
// 		go func() {
// 			fmt.Printf("Test Channel...%d\n", i)
// 			tc <- 1 //  say "finish!"
// 			// tc <- 1 //  say "finish!"
// 		}()

// 		fmt.Printf("No Channel...%d\n", i)
// 		<-tc // wait
// 	}

// }

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
	var tmp int = vec.Get(0).(int) // force type conversion
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
		//fmt.Printf("Now The Size : %d,  Now The Capacity: %d \n", vec.Size(), vec.Capacity())
	}

}
