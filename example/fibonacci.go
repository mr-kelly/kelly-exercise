package example

import (
// "fmt"
)

// 传统数学定义上, 递归
func fib2(n int) uint {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

// 递推法
func fib(n int) uint {
	tmpList := make([]uint, n+1)
	if n == 0 {
		return 1
	}

	for i := 0; i <= n; i++ {
		if i < 2 {
			tmpList[i] = 1 // result 1   n 为2时
		} else {
			tmpList[i] = tmpList[i-1] + tmpList[i-2]
		}

	}
	return tmpList[n]
}

// func main() {
// 	fmt.Printf("Fib Result: %d\n", fib(1))
// 	fmt.Printf("Fib Result: %d\n", fib2(1))

// 	fmt.Printf("Fib Result: %d\n", fib(2))
// 	fmt.Printf("Fib Result: %d\n", fib2(2))

// 	fmt.Printf("Fib Result: %d\n", fib(3))
// 	fmt.Printf("Fib Result: %d\n", fib2(3))

// 	fmt.Printf("Fib Result: %d\n", fib(4))
// 	fmt.Printf("Fib Result: %d\n", fib2(4))

// 	fmt.Printf("Fib Result: %d\n", fib2(100))
// 	fmt.Printf("Fib Result: %d\n", fib(100))
// }
