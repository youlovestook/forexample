package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("cpu: ", runtime.NumCPU())
	fmt.Println("Go Routines: ", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < 100; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Go Routines each loop: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("counter:", counter)
}
