package example

import (
	// "fmt"
	"testing"
)

func TestFib(t *testing.T) {
	if fib(100) != 2425370821 {
		t.Errorf("Error Fib")
	}

}

func TestFib2(t *testing.T) {
	if fib2(100) != 2425370821 {
		t.Errorf("Error Fib2")
	}
}

// 性能测试
func BenchmarkFib2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib2(100)
	}

}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(100)
	}
}
