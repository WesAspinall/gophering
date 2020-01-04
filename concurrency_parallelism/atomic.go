package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// using atomic to add a value to an int64
// load value from int64 -- basically reading it
func main() {
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)                         // write to counter
			fmt.Println("Counter\t", atomic.LoadInt64(&counter)) //read from counter
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
