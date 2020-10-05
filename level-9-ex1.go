package main

import (
	"fmt"
	"sync"
)

// concurrency is a design pattern
// go is the first language writted to take advantage of
// computers with multiple cores
// rob pike "concurrency is not parrallelism " on youtube

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go foo()

	go bar()

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
	wg.Done()
}
