package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("cpu: ", runtime.NumCPU())
	fmt.Println("Go Routines: ", runtime.NumGoroutine())
	var mu sync.Mutex
	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Go Routines each loop: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("counter:", counter)
}
