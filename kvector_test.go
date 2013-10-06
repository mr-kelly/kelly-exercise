package klib

import (
	"testing"
)

func TestKVector(t *testing.T) {

	// Test Pointer type
	var vec *KVector = NewKVector()
	var tmp *KVector_T = new(KVector_T)
	var tmp2 int = 2
	//vec.PushBack(tmp)
	vec.PushBack((*KVector_T(tmp2)))
	if vec.Size() != 1 {
		//t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
		t.Errorf("vec.Size() != 1")
	}

	// Test Value type
	// var vec *KVector = NewKVector()
	// var tmp *KVector_T = new(KVector_T)

	// vec.PushBack(tmp)
	// if vec.Size() != 1 {
	// 	//t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
	// 	t.Errorf("vec.Size() != 1")
	// }
}
