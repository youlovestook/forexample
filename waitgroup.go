package main

import (
	"fmt"
	"runtime"
	"sync"
)

// concurrency is a design pattern
// go is the first language writted to take advantage of
// computers with multiple cores
// rob pike "concurrency is not parrallelism " on youtube

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)

	go foo()
	fmt.Println("------------------------")
	bar()
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()

}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo: ", i)
	}
	wg.Done()
}

func bar() {
	for j := 0; j < 10; j++ {
		fmt.Println("bar: ", j)
	}
}
