package main

import (
	"fmt"
	"os"
	"sync"
)

var test int = 3
var testChan = make(chan int, 0)

var l *sync.Mutex
var a string = ""
var once sync.Once

// var done bool

var myStdout = os.Stdout

func print(str string) {
	fmt.Printf(str)
}

func main() {
	twoprint()
}

func setup() {
	a = "hello, world"

	test++
	// done = true
}

func doprint() {
	os.Stdout = myStdout
	once.Do(setup) // only one in goroutine

	testChan <- 1
}

func twoprint() {
	go doprint()
	go doprint()

	<-testChan
	fmt.Printf("%d", test)

}
