package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("cpu: ", runtime.NumCPU())
	fmt.Println("Go Routines: ", runtime.NumGoroutine())
	//var mu sync.Mutex
	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < 100; i++ {
		go func() {

			atomic.AddInt64(&counter, 1)
			fmt.Println("counter:   ", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("Go Routines each loop: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("counter:", counter)
}
